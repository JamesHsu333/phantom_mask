package service

import (
	"time"

	"github.com/JamesHsu333/kdan/config"
	"github.com/JamesHsu333/kdan/pkg/logger"
	"github.com/jackc/pgx/v5"
)

const (
	waitShotDownDuration    = 3 * time.Second
	initPostgresSQLAttempts = 5
	initPostgresSQLDelay    = time.Duration(1500) * time.Millisecond
)

// Service
type Service struct {
	cfg    *config.Config
	db     *pgx.Conn
	logger logger.Logger
	doneCh chan struct{}
}
