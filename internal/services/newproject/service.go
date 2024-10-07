package newproject

import (
	"github.com/upikoth/starter-new/internal/config"
	"github.com/upikoth/starter-new/internal/model"
	"github.com/upikoth/starter-new/internal/pkg/logger"
	"github.com/upikoth/starter-new/internal/repositories"
	"github.com/upikoth/starter-new/internal/services/ycuser"
)

type Service struct {
	newProject    *model.Project
	logger        logger.Logger
	config        *config.Config
	repositories  *repositories.Repositories
	ycUserService *ycuser.Service
}

func New(
	logger logger.Logger,
	config *config.Config,
	repositories *repositories.Repositories,
	ycUserService *ycuser.Service,
) *Service {

	return &Service{
		newProject: model.NewProject(
			config,
		),
		logger:        logger,
		config:        config,
		repositories:  repositories,
		ycUserService: ycUserService,
	}
}
