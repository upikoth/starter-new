package app

import (
	"context"

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

func New(
	config *config.Config,
	logger logger.Logger,
) (*App, error) {
	repositoriesInstance, err := repositories.New(logger, config)

	if err != nil {
		return nil, err
	}

	servicesInstance, err := services.New(logger, config, repositoriesInstance)

	if err != nil {
		return nil, err
	}

	return &App{
		config:       config,
		logger:       logger,
		repositories: repositoriesInstance,
		services:     servicesInstance,
	}, nil
}

func (s *App) Start(ctx context.Context) error {
	s.logger.Info("Задаем имя проекта")
	err := s.services.NewProjectService.CreateNewProjectName(ctx)

	if err != nil {
		return err
	}
	s.logger.Info("Имя проекта успешно задано")

	s.logger.Info(`Шаг 1: Создаем
			github репозиторий,
			folder в yandex.cloud
		`)
	eg, newCtx := errgroup.WithContext(ctx)

	//eg.Go(func() error {
	//	return s.services.NewProjectService.CreateGithubRepositories(newCtx)
	//})

	eg.Go(func() error {
		return s.services.NewProjectService.CreateYCFolder(newCtx)
	})

	err = eg.Wait()

	if err != nil {
		return err
	}
	s.logger.Info("Шаг 1 успешно выполнен!")

	s.logger.Info(`Шаг 2: Создаем 
			сервисный аккаунт,
			бакеты,
			container registry,
			ydb,
			serverless container,
			logging group,
			dns zone
		`)
	eg, newCtx = errgroup.WithContext(ctx)

	eg.Go(func() error {
		return s.services.NewProjectService.CreateYCFolderServiceAccount(newCtx)
	})

	eg.Go(func() error {
		return s.services.NewProjectService.CreateYCStorageBuckets(newCtx)
	})

	eg.Go(func() error {
		return s.services.NewProjectService.CreateYCContainerRegistry(newCtx)
	})

	eg.Go(func() error {
		return s.services.NewProjectService.CreateYCYDB(newCtx)
	})

	eg.Go(func() error {
		return s.services.NewProjectService.CreateYCServerlessContainer(newCtx)
	})

	eg.Go(func() error {
		return s.services.NewProjectService.CreateYCLogGroup(newCtx)
	})

	eg.Go(func() error {
		return s.services.NewProjectService.CreateYCDNSZone(newCtx)
	})

	eg.Go(func() error {
		return s.services.NewProjectService.CreateYCCertificate(ctx)
	})

	err = eg.Wait()

	if err != nil {
		return err
	}
	s.logger.Info("Шаг 2 успешно выполнен!")

	s.logger.Info(`Шаг 3: Создаем 
			api gateway
			postbox address
			запись о сертификате в dns
			dns record для подтверждения, что почта относится к нашему домену
		`)
	eg, newCtx = errgroup.WithContext(ctx)

	eg.Go(func() error {
		return s.services.NewProjectService.CreateYCApiGateway(ctx)
	})

	eg.Go(func() error {
		return s.services.NewProjectService.CreateYCPostboxAddress(ctx)
	})

	eg.Go(func() error {
		return s.services.NewProjectService.BindCertificateToDNSZone(ctx)
	})

	err = eg.Wait()

	if err != nil {
		return err
	}
	s.logger.Info("Шаг 3 успешно выполнен!")

	s.logger.Info(`Шаг 4: Создаем 
			dns record для подтверждения, что почта относится к нашему домену
			привязку gateway к домену
		`)
	eg, newCtx = errgroup.WithContext(ctx)

	eg.Go(func() error {
		return s.services.NewProjectService.AddYCPostboxDNSRecord(ctx)
	})

	eg.Go(func() error {
		return s.services.NewProjectService.BindYCGatewayToDNS(ctx)
	})

	err = eg.Wait()

	if err != nil {
		return err
	}
	s.logger.Info("Шаг 4 успешно выполнен!")

	return nil
}

func (s *App) Stop(_ context.Context) error {
	return nil
}
