package repository

import (
	"github.com/upikoth/starter-new/internal/config"
	"github.com/upikoth/starter-new/internal/pkg/logger"
	"github.com/upikoth/starter-new/internal/repository/consoleinput"
	"github.com/upikoth/starter-new/internal/repository/github"
)

type Repository struct {
	ConsoleInput *consoleinput.ConsoleInput
	Github       *github.Github
}

func New(
	logger logger.Logger,
	config *config.Config,
) (*Repository, error) {
	return &Repository{
		ConsoleInput: consoleinput.New(
			logger,
			config,
		),
		Github: github.New(
			logger,
			config,
		),
	}, nil
}
