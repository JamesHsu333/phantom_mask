package postgres

import (
	"context"
	"fmt"

	"github.com/jackc/pgx/v5"
)

type Config struct {
	PostgresqlHost     string
	PostgresqlPort     string
	PostgresqlUser     string
	PostgresqlPassword string
	PostgresqlDbname   string
	PostgresqlSSLMode  bool
	PgDriver           string
}

// Return new Postgresql db instance
func NewPsqlDB(ctx context.Context, cfg Config) (*pgx.Conn, error) {
	connString := fmt.Sprintf("postgres://%s:%s@%s:%s/%s",
		cfg.PostgresqlUser,
		cfg.PostgresqlPassword,
		cfg.PostgresqlHost,
		cfg.PostgresqlPort,
		cfg.PostgresqlDbname,
	)

	db, err := pgx.Connect(context.Background(), connString)
	if err != nil {
		return nil, err
	}

	if err = db.Ping(ctx); err != nil {
		return nil, err
	}

	return db, nil
}
