package app

import (
	"context"
	"fmt"

	"github.com/upikoth/starter-new/internal/config"
	"github.com/upikoth/starter-new/internal/pkg/logger"
	"github.com/upikoth/starter-new/internal/repositories"
	"github.com/upikoth/starter-new/internal/services"
	"golang.org/x/sync/errgroup"
)

type App struct {
	config       *config.Config
	logger       logger.Logger
	repositories *repositories.Repositories
	services     *services.Services
}

func New(config *config.Config, logger logger.Logger) (*App, error) {
	repositories, err := repositories.New(logger, config)

	if err != nil {
		logger.Error(fmt.Sprintf("Ошибка при инициализации repository: %s", err))
		return nil, err
	}

	services, err := services.New(logger, config, repositories)

	if err != nil {
		logger.Error(fmt.Sprintf("Ошибка при инициализации service: %s", err))
		return nil, err
	}

	return &App{
		config:       config,
		logger:       logger,
		repositories: repositories,
		services:     services,
	}, nil
}

func (s *App) Start(ctx context.Context) error {
	if err := s.services.NewProject.CreateNewProjectName(ctx); err != nil {
		return err
	}

	// 1. Создаем:
	// github репозитории
	// folder в yandex.cloud
	eg, newCtx := errgroup.WithContext(ctx)

	eg.Go(func() error {
		return s.services.NewProject.CreateGithubRepositories(newCtx)
	})

	eg.Go(func() error {
		return s.services.NewProject.CreateYCFolder(newCtx)
	})

	err := eg.Wait()

	if err != nil {
		return err
	}

	eg, newCtx = errgroup.WithContext(ctx)

	// 2. Создаем:
	// сервисный аккаунт
	// бакеты
	eg.Go(func() error {
		return s.services.NewProject.CreateYCFolderServiceAccount(newCtx)
	})

	eg.Go(func() error {
		return s.services.NewProject.CreateYCStorageBuckets(newCtx)
	})

	eg.Go(func() error {
		return s.services.NewProject.CreateYCContainerRegistry(newCtx)
	})

	eg.Go(func() error {
		return s.services.NewProject.CreateYCYDB(newCtx)
	})

	err = eg.Wait()

	if err != nil {
		return err
	}

	return nil
}

func (s *App) Stop(_ context.Context) error {
	return nil
}
