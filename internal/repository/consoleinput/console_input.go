package consoleinput

import (
	"github.com/upikoth/starter-new/internal/config"
	"github.com/upikoth/starter-new/internal/pkg/logger"
)

type ConsoleInput struct {
	logger logger.Logger
	config *config.Config
}

func New(
	logger logger.Logger,
	config *config.Config,
) *ConsoleInput {
	return &ConsoleInput{
		logger: logger,
		config: config,
	}
}
