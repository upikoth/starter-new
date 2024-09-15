package service

import (
	"github.com/upikoth/starter-new/internal/config"
	"github.com/upikoth/starter-new/internal/pkg/logger"
	"github.com/upikoth/starter-new/internal/repository"
	"github.com/upikoth/starter-new/internal/service/newproject"
)

type Service struct {
	NewProject *newproject.NewProject
}

func New(
	logger logger.Logger,
	config *config.Config,
	repository *repository.Repository,
) (*Service, error) {
	return &Service{
		NewProject: newproject.New(
			logger,
			config,
			repository,
		),
	}, nil
}
