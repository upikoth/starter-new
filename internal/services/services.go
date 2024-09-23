package services

import (
	"github.com/upikoth/starter-new/internal/config"
	"github.com/upikoth/starter-new/internal/pkg/logger"
	"github.com/upikoth/starter-new/internal/repositories"
	"github.com/upikoth/starter-new/internal/services/newproject"
	"github.com/upikoth/starter-new/internal/services/user"
)

type Services struct {
	NewProject *newproject.NewProject
	User       *user.User
}

func New(
	logger logger.Logger,
	config *config.Config,
	repositories *repositories.Repositories,
) (*Services, error) {
	return &Services{
		User: user.New(
			logger,
			config,
			repositories,
		),
		NewProject: newproject.New(
			logger,
			config,
			repositories,
		),
	}, nil
}
