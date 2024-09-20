package services

import (
	"github.com/upikoth/starter-new/internal/config"
	"github.com/upikoth/starter-new/internal/pkg/logger"
	"github.com/upikoth/starter-new/internal/repositories"
	"github.com/upikoth/starter-new/internal/services/newproject"
)

type Services struct {
	NewProject *newproject.NewProject
}

func New(
	logger logger.Logger,
	config *config.Config,
	repositories *repositories.Repositories,
) (*Services, error) {
	return &Services{
		NewProject: newproject.New(
			logger,
			config,
			repositories,
		),
	}, nil
}
