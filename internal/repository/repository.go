package repository

import (
	"github.com/upikoth/starter-new/internal/config"
	"github.com/upikoth/starter-new/internal/pkg/logger"
	"github.com/upikoth/starter-new/internal/repository/consoleinput"
)

type Repository struct {
	ConsoleInput consoleinput.ConsoleInput
}

func New(
	logger logger.Logger,
	config *config.Config,
) (*Repository, error) {
	return &Repository{
		ConsoleInput: *consoleinput.New(
			logger,
			config,
		),
	}, nil
}
