package newproject

import (
	"github.com/upikoth/starter-new/internal/config"
	"github.com/upikoth/starter-new/internal/model"
	"github.com/upikoth/starter-new/internal/pkg/logger"
	"github.com/upikoth/starter-new/internal/repository"
)

type NewProject struct {
	project    *model.Project
	logger     logger.Logger
	config     *config.Config
	repository *repository.Repository
}

func New(
	logger logger.Logger,
	config *config.Config,
	repository *repository.Repository,
) *NewProject {
	return &NewProject{
		logger:     logger,
		config:     config,
		repository: repository,
		project:    &model.Project{},
	}
}
