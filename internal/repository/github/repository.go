package github

import (
	"github.com/upikoth/starter-new/internal/config"
	"github.com/upikoth/starter-new/internal/pkg/logger"
)

type Github struct {
	logger logger.Logger
	config *config.Config
}

func New(
	logger logger.Logger,
	config *config.Config,
) *Github {
	return &Github{
		logger: logger,
		config: config,
	}
}
