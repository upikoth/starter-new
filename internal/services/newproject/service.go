package newproject

import (
	"github.com/upikoth/starter-new/internal/config"
	"github.com/upikoth/starter-new/internal/model"
	"github.com/upikoth/starter-new/internal/pkg/logger"
	"github.com/upikoth/starter-new/internal/repositories"
)

type NewProject struct {
	project      *model.Project
	logger       logger.Logger
	config       *config.Config
	repositories *repositories.Repositories
}

func New(
	logger logger.Logger,
	config *config.Config,
	repositories *repositories.Repositories,
) *NewProject {
	return &NewProject{
		logger:       logger,
		config:       config,
		repositories: repositories,
		project:      &model.Project{},
	}
}
