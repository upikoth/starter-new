package services

import (
	"github.com/upikoth/starter-new/internal/config"
	"github.com/upikoth/starter-new/internal/pkg/logger"
	"github.com/upikoth/starter-new/internal/repositories"
	"github.com/upikoth/starter-new/internal/services/newproject"
	"github.com/upikoth/starter-new/internal/services/ycuser"
)

type Services struct {
	NewProjectService *newproject.NewProjectService
	YCUserService     *ycuser.YCUserService
}

func New(
	logger logger.Logger,
	config *config.Config,
	repositories *repositories.Repositories,
) (*Services, error) {
	ycUserService := ycuser.New(
		logger,
		config,
		repositories,
	)

	newProjectService := newproject.New(
		logger,
		config,
		repositories,
		ycUserService,
	)

	return &Services{
		YCUserService:     ycUserService,
		NewProjectService: newProjectService,
	}, nil
}
