package service

import (
	"context"
	"fmt"
	"net"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/alexliesenfeld/health"
	"github.com/avast/retry-go"
	grpcrecovery "github.com/grpc-ecosystem/go-grpc-middleware/recovery"
	grpc_ctxtags "github.com/grpc-ecosystem/go-grpc-middleware/tags"
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"github.com/pkg/errors"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/keepalive"
	"google.golang.org/grpc/reflection"

	"github.com/JamesHsu333/kdan/config"
	data "github.com/JamesHsu333/kdan/internal/data/test"
	"github.com/JamesHsu333/kdan/internal/interceptors"
	"github.com/JamesHsu333/kdan/internal/service/pharmacy"
	"github.com/JamesHsu333/kdan/internal/service/user"
	"github.com/JamesHsu333/kdan/pkg/database/postgres"
	"github.com/JamesHsu333/kdan/pkg/logger"
	kdanProto "github.com/JamesHsu333/kdan/proto/kdan"
)

// NewService New Service constructor
func NewService(cfg *config.Config, logger logger.Logger) *Service {
	return &Service{cfg: cfg, logger: logger, doneCh: make(chan struct{})}
}

func (s *Service) Run() error {
	ctx, cancel := signal.NotifyContext(context.Background(), os.Interrupt, syscall.SIGTERM, syscall.SIGINT)
	defer cancel()

	err := s.init(ctx)
	if err != nil {
		return err
	}
	defer s.closeConn(context.Background())

	grpcServer := s.createGRPCServer()

	l, err := net.Listen("tcp", s.cfg.Server.Port)
	if err != nil {
		return err
	}
	defer l.Close()

	httpServer, err := s.createHTTPServer(ctx)
	if err != nil {
		return err
	}

	healthCheckServer := s.createHealthCheckServer()

	go func() {
		s.logger.Infof("Server is listening on port: %v", s.cfg.Server.Port)
		if err := grpcServer.Serve(l); err != nil {
			s.logger.Error(err)
		}
	}()

	go func() {
		s.logger.Infof("HTTP Server is listening on port: %v", s.cfg.Server.HttpPort)
		if err := httpServer.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			s.logger.Error(err)
		}
	}()

	go func() {
		s.logger.Infof("Health Check Server is listening on: %v%v", s.cfg.HealthCheck.URL, s.cfg.HealthCheck.Path)
		if err := healthCheckServer.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			s.logger.Error(err)
		}
	}()

	<-ctx.Done()
	s.waitShotDown(waitShotDownDuration)

	grpcServer.GracefulStop()

	if err := httpServer.Shutdown(ctx); err != nil && err != http.ErrServerClosed {
		s.logger.Warnf("(Shutdown) httpServer err: %v", err)
	}

	if err := healthCheckServer.Shutdown(ctx); err != nil && err != http.ErrServerClosed {
		s.logger.Warnf("(Shutdown) healthCheckServer err: %v", err)
	}

	<-s.doneCh
	s.logger.Info("Server Exited Properly")
	return nil
}

func (s *Service) init(ctx context.Context) error {
	if err := s.initPostgreSQL(ctx); err != nil {
		return err
	}

	return nil
}

func (s *Service) closeConn(ctx context.Context) {
	s.closePostgresConn(ctx)
}

func (s *Service) createGRPCServer() *grpc.Server {
	querier := data.New(s.db)

	userUC := user.NewUserUC(querier, s.db, s.logger)
	pharmacyUC := pharmacy.NewPharmacyUC(querier, s.logger)
	im := interceptors.NewInterceptorManager(s.logger, s.cfg)

	service := &kdanService{}

	service.UserEndpoint = *user.NewUserEndpoint(s.logger, s.cfg, userUC)
	service.PharmacyEndpoint = *pharmacy.NewPharmacyEndpoint(s.logger, s.cfg, pharmacyUC)

	server := grpc.NewServer(grpc.KeepaliveParams(keepalive.ServerParameters{
		MaxConnectionIdle: s.cfg.Server.MaxConnectionIdle * time.Minute,
		Timeout:           s.cfg.Server.Timeout * time.Second,
		MaxConnectionAge:  s.cfg.Server.MaxConnectionAge * time.Minute,
		Time:              s.cfg.Server.Timeout * time.Minute,
	}),
		grpc.ChainUnaryInterceptor(
			im.Logger,
			grpc_ctxtags.UnaryServerInterceptor(),
			grpcrecovery.UnaryServerInterceptor(),
		),
	)

	if s.cfg.Server.Mode != "Production" {
		reflection.Register(server)
	}

	kdanProto.RegisterKdanServiceServer(server, service)

	return server
}

func (s *Service) createHTTPServer(ctx context.Context) (*http.Server, error) {
	mux := runtime.NewServeMux()
	opts := []grpc.DialOption{grpc.WithTransportCredentials(insecure.NewCredentials())}
	grpcServerEndpoint := fmt.Sprintf("localhost%s", s.cfg.Server.Port)
	err := kdanProto.RegisterKdanServiceHandlerFromEndpoint(ctx, mux, grpcServerEndpoint, opts)
	if err != nil {
		return nil, err
	}

	return &http.Server{
		Addr:    s.cfg.Server.HttpPort,
		Handler: mux,
	}, err
}

func (s *Service) createHealthCheckServer() *http.Server {
	checker := health.NewChecker(
		health.WithCacheDuration(1*time.Second),
		health.WithTimeout(10*time.Second),
		health.WithPeriodicCheck(15*time.Second, 3*time.Second, health.Check{
			Name:    "database",
			Timeout: 2 * time.Second,
			Check:   s.db.Ping,
		}),
		health.WithStatusListener(func(ctx context.Context, state health.CheckerState) {
			switch state.Status {
			case health.StatusUp:
				s.logger.Infof("Overall system health status changed to %s", state.Status)
			case health.StatusDown:
				s.logger.Warnf("Overall system health status changed to %s", state.Status)
				// TODO: recover connection
			default:
				s.logger.Warnf("Overall system health status changed to %s", state.Status)
				// TODO: recover connection
			}
		}),
	)

	mux := http.NewServeMux()
	mux.Handle(s.cfg.HealthCheck.Path, health.NewHandler(checker))
	return &http.Server{
		Addr:    s.cfg.HealthCheck.URL,
		Handler: mux,
	}
}

func (s *Service) initPostgreSQL(ctx context.Context) error {
	retryOptions := []retry.Option{
		retry.Attempts(initPostgresSQLAttempts),
		retry.Delay(initPostgresSQLDelay),
		retry.DelayType(retry.BackOffDelay),
		retry.LastErrorOnly(true),
		retry.Context(ctx),
		retry.OnRetry(func(n uint, err error) {
			s.logger.Errorf("retry connect postgresql err: %v", err)
		}),
	}

	return retry.Do(func() (err error) {
		s.db, err = postgres.NewPsqlDB(ctx, s.cfg.Postgres)
		if err != nil {
			return errors.Wrap(err, "Postgresql init")
		}

		s.logger.Infof("Postgres connected")
		return nil
	}, retryOptions...)

}

func (s *Service) closePostgresConn(ctx context.Context) {
	if err := s.db.Close(ctx); err != nil {
		s.logger.Errorf("s.db.Close err: %v", err)
	}
}

func (s *Service) waitShotDown(duration time.Duration) {
	go func() {
		time.Sleep(duration)
		s.doneCh <- struct{}{}
	}()
}
