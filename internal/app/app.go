package app

import (
	"context"
	"fmt"

	"github.com/upikoth/starter-new/internal/config"
	"github.com/upikoth/starter-new/internal/pkg/logger"
	"github.com/upikoth/starter-new/internal/repository"
	"github.com/upikoth/starter-new/internal/service"
)

type App struct {
	config     *config.Config
	logger     logger.Logger
	repository *repository.Repository
	service    *service.Service
}

func New(config *config.Config, logger logger.Logger) (*App, error) {
	repository, err := repository.New(logger, config)

	if err != nil {
		logger.Error(fmt.Sprintf("Ошибка при инициализации repository: %s", err))
		return nil, err
	}

	service, err := service.New(logger, config, repository)

	if err != nil {
		logger.Error(fmt.Sprintf("Ошибка при инициализации service: %s", err))
		return nil, err
	}

	return &App{
		config:     config,
		logger:     logger,
		repository: repository,
		service:    service,
	}, nil
}

func (s *App) Start() error {
	if err := s.service.NewProject.SetNewProjectName(); err != nil {
		return err
	}

	s.service.NewProject.CreateGithubRepositories()

	return nil
}

func (s *App) Stop(_ context.Context) error {
	return nil
}
