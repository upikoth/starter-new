package fileinput

import (
	"github.com/upikoth/starter-new/internal/config"
	"github.com/upikoth/starter-new/internal/pkg/logger"
)

type FileInput struct {
	logger logger.Logger
	config *config.Config
}

func New(
	logger logger.Logger,
	config *config.Config,
) *FileInput {
	return &FileInput{
		logger: logger,
		config: config,
	}
}
