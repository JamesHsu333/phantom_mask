package pharmacy

import (
	"github.com/JamesHsu333/kdan/config"
	"github.com/JamesHsu333/kdan/pkg/logger"
)

type PharmacyEndpoint struct {
	logger     logger.Logger
	cfg        *config.Config
	pharmacyUC *pharmacyUC
}
