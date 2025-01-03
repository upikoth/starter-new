package app

import (
	"context"
	"github.com/upikoth/starter-new/internal/pkg/functionswithneeds"
	"github.com/upikoth/starter-new/internal/pkg/logger"

	"github.com/upikoth/starter-new/internal/config"
	"github.com/upikoth/starter-new/internal/repositories"
	"github.com/upikoth/starter-new/internal/services"
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
	return functionswithneeds.Start(
		ctx,
		functionswithneeds.FunctionsWithNeeds{
			functionswithneeds.FunctionWithNeeds{
				Function: s.services.NewProjectService.CreateNewProjectName,
				Needs:    nil,
			},
			functionswithneeds.FunctionWithNeeds{
				Function: s.services.NewProjectService.CreateSentryBackendProject,
				Needs: []func(ctx context.Context) error{
					s.services.NewProjectService.CreateNewProjectName,
				},
			},
			functionswithneeds.FunctionWithNeeds{
				Function: s.services.NewProjectService.CreateSentryFrontendProject,
				Needs: []func(ctx context.Context) error{
					s.services.NewProjectService.CreateNewProjectName,
				},
			},
			functionswithneeds.FunctionWithNeeds{
				Function: s.services.NewProjectService.CreateSentryBackendDSN,
				Needs: []func(ctx context.Context) error{
					s.services.NewProjectService.CreateSentryBackendProject,
				},
			},
			functionswithneeds.FunctionWithNeeds{
				Function: s.services.NewProjectService.CreateSentryFrontendDSN,
				Needs: []func(ctx context.Context) error{
					s.services.NewProjectService.CreateSentryFrontendProject,
				},
			},
			functionswithneeds.FunctionWithNeeds{
				Function: s.services.NewProjectService.CreateLocalRepos,
				Needs: []func(ctx context.Context) error{
					s.services.NewProjectService.CreateNewProjectName,
				},
			},
			functionswithneeds.FunctionWithNeeds{
				Function: s.services.NewProjectService.CreateGithubRepositories,
				Needs: []func(ctx context.Context) error{
					s.services.NewProjectService.CreateNewProjectName,
				},
			},
			functionswithneeds.FunctionWithNeeds{
				Function: s.services.NewProjectService.CreateYCFolder,
				Needs: []func(ctx context.Context) error{
					s.services.NewProjectService.CreateNewProjectName,
				},
			},
			functionswithneeds.FunctionWithNeeds{
				Function: s.services.NewProjectService.AddRenovateToGithubRepositories,
				Needs: []func(ctx context.Context) error{
					s.services.NewProjectService.CreateGithubRepositories,
				},
			},
			functionswithneeds.FunctionWithNeeds{
				Function: s.services.NewProjectService.SetGithubRepositoriesHomepage,
				Needs: []func(ctx context.Context) error{
					s.services.NewProjectService.CreateGithubRepositories,
				},
			},
			functionswithneeds.FunctionWithNeeds{
				Function: s.services.NewProjectService.CreateYCFolderServiceAccount,
				Needs: []func(ctx context.Context) error{
					s.services.NewProjectService.CreateYCFolder,
				},
			},
			functionswithneeds.FunctionWithNeeds{
				Function: s.services.NewProjectService.CreateYCStorageBuckets,
				Needs: []func(ctx context.Context) error{
					s.services.NewProjectService.CreateYCFolder,
				},
			},
			functionswithneeds.FunctionWithNeeds{
				Function: s.services.NewProjectService.CreateYCContainerRegistry,
				Needs: []func(ctx context.Context) error{
					s.services.NewProjectService.CreateYCFolder,
				},
			},
			functionswithneeds.FunctionWithNeeds{
				Function: s.services.NewProjectService.CreateYCYDB,
				Needs: []func(ctx context.Context) error{
					s.services.NewProjectService.CreateYCFolder,
				},
			},
			functionswithneeds.FunctionWithNeeds{
				Function: s.services.NewProjectService.CreateYCServerlessContainer,
				Needs: []func(ctx context.Context) error{
					s.services.NewProjectService.CreateYCFolder,
				},
			},
			functionswithneeds.FunctionWithNeeds{
				Function: s.services.NewProjectService.CreateYCLogGroup,
				Needs: []func(ctx context.Context) error{
					s.services.NewProjectService.CreateYCFolder,
				},
			},
			functionswithneeds.FunctionWithNeeds{
				Function: s.services.NewProjectService.CreateYCDNSZone,
				Needs: []func(ctx context.Context) error{
					s.services.NewProjectService.CreateYCFolder,
				},
			},
			functionswithneeds.FunctionWithNeeds{
				Function: s.services.NewProjectService.CreateYCCertificate,
				Needs: []func(ctx context.Context) error{
					s.services.NewProjectService.CreateYCFolder,
				},
			},
			functionswithneeds.FunctionWithNeeds{
				Function: s.services.NewProjectService.YCCreateObjectStorageAccessKey,
				Needs: []func(ctx context.Context) error{
					s.services.NewProjectService.CreateYCFolderServiceAccount,
				},
			},
			functionswithneeds.FunctionWithNeeds{
				Function: s.services.NewProjectService.YCCreatePostboxAccessKey,
				Needs: []func(ctx context.Context) error{
					s.services.NewProjectService.CreateYCFolderServiceAccount,
				},
			},
			functionswithneeds.FunctionWithNeeds{
				Function: s.services.NewProjectService.YCCreateServiceAccountAuthorizedKey,
				Needs: []func(ctx context.Context) error{
					s.services.NewProjectService.CreateYCFolderServiceAccount,
				},
			},
			functionswithneeds.FunctionWithNeeds{
				Function: s.services.NewProjectService.UploadYCSecretsToObjectStorage,
				Needs: []func(ctx context.Context) error{
					s.services.NewProjectService.YCCreateServiceAccountAuthorizedKey,
					s.services.NewProjectService.CreateYCStorageBuckets,
				},
			},
			functionswithneeds.FunctionWithNeeds{
				Function: s.services.NewProjectService.UpdateYCAccessToRegistry,
				Needs: []func(ctx context.Context) error{
					s.services.NewProjectService.CreateYCContainerRegistry,
					s.services.NewProjectService.CreateYCFolderServiceAccount,
				},
			},
			functionswithneeds.FunctionWithNeeds{
				Function: s.services.NewProjectService.CreateYCApiGateway,
				Needs: []func(ctx context.Context) error{
					s.services.NewProjectService.CreateYCFolder,
					s.services.NewProjectService.CreateYCLogGroup,
					s.services.NewProjectService.CreateYCFolderServiceAccount,
					s.services.NewProjectService.CreateYCServerlessContainer,
				},
			},
			functionswithneeds.FunctionWithNeeds{
				Function: s.services.NewProjectService.CreateYCPostboxAddress,
				Needs: []func(ctx context.Context) error{
					s.services.NewProjectService.CreateYCFolder,
					s.services.NewProjectService.CreateYCLogGroup,
				},
			},
			functionswithneeds.FunctionWithNeeds{
				Function: s.services.NewProjectService.BindCertificateToDNSZone,
				Needs: []func(ctx context.Context) error{
					s.services.NewProjectService.CreateYCFolder,
					s.services.NewProjectService.CreateYCCertificate,
					s.services.NewProjectService.CreateYCDNSZone,
				},
			},
			functionswithneeds.FunctionWithNeeds{
				Function: s.services.NewProjectService.SetupGithubBackendRepo,
				Needs: []func(ctx context.Context) error{
					s.services.NewProjectService.CreateYCFolder,
					s.services.NewProjectService.CreateYCLogGroup,
					s.services.NewProjectService.CreateYCFolderServiceAccount,
					s.services.NewProjectService.CreateYCContainerRegistry,
					s.services.NewProjectService.CreateLocalRepos,
					s.services.NewProjectService.CreateGithubRepositories,
					s.services.NewProjectService.YCCreatePostboxAccessKey,
					s.services.NewProjectService.CreateYCYDB,
					s.services.NewProjectService.YCCreateServiceAccountAuthorizedKey,
					s.services.NewProjectService.UpdateYCAccessToRegistry,
					s.services.NewProjectService.CreateSentryBackendDSN,
				},
			},
			functionswithneeds.FunctionWithNeeds{
				Function: s.services.NewProjectService.CreateBackendLocalDevEnvironment,
				Needs: []func(ctx context.Context) error{
					s.services.NewProjectService.CreateYCFolder,
					s.services.NewProjectService.CreateYCLogGroup,
					s.services.NewProjectService.CreateYCFolderServiceAccount,
					s.services.NewProjectService.CreateYCContainerRegistry,
					s.services.NewProjectService.CreateLocalRepos,
					s.services.NewProjectService.YCCreatePostboxAccessKey,
					s.services.NewProjectService.CreateYCYDB,
					s.services.NewProjectService.YCCreateServiceAccountAuthorizedKey,
					s.services.NewProjectService.UpdateYCAccessToRegistry,
					s.services.NewProjectService.CreateSentryBackendDSN,
				},
			},
			functionswithneeds.FunctionWithNeeds{
				Function: s.services.NewProjectService.SetupGithubFrontendRepo,
				Needs: []func(ctx context.Context) error{
					s.services.NewProjectService.CreateLocalRepos,
					s.services.NewProjectService.CreateGithubRepositories,
					s.services.NewProjectService.YCCreateObjectStorageAccessKey,
					s.services.NewProjectService.CreateSentryFrontendDSN,
				},
			},
			functionswithneeds.FunctionWithNeeds{
				Function: s.services.NewProjectService.CreateFrontendLocalDevEnvironment,
				Needs: []func(ctx context.Context) error{
					s.services.NewProjectService.CreateLocalRepos,
					s.services.NewProjectService.YCCreateObjectStorageAccessKey,
					s.services.NewProjectService.CreateSentryFrontendDSN,
				},
			},
			functionswithneeds.FunctionWithNeeds{
				Function: s.services.NewProjectService.AddYCPostboxDNSRecord,
				Needs: []func(ctx context.Context) error{
					s.services.NewProjectService.CreateYCPostboxAddress,
					s.services.NewProjectService.CreateYCDNSZone,
				},
			},
			functionswithneeds.FunctionWithNeeds{
				Function: s.services.NewProjectService.BindYCGatewayToDNS,
				Needs: []func(ctx context.Context) error{
					s.services.NewProjectService.CreateYCDNSZone,
					s.services.NewProjectService.CreateYCApiGateway,
					s.services.NewProjectService.CreateYCCertificate,
				},
			},
		},
	)
}

func (s *App) Stop(_ context.Context) error {
	return nil
}
