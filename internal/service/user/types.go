package user

import (
	"github.com/JamesHsu333/kdan/config"
	"github.com/JamesHsu333/kdan/pkg/logger"
)

type UserEndpoint struct {
	logger logger.Logger
	cfg    *config.Config
	userUC *userUC
}
