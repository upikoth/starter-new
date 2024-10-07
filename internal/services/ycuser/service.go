package ycuser

import (
	"github.com/upikoth/starter-new/internal/config"
	"github.com/upikoth/starter-new/internal/model"
	"github.com/upikoth/starter-new/internal/pkg/logger"
	"github.com/upikoth/starter-new/internal/repositories"
)

type Service struct {
	ycUser       *model.YCUser
	logger       logger.Logger
	config       *config.Config
	repositories *repositories.Repositories
}

func New(
	logger logger.Logger,
	config *config.Config,
	repositories *repositories.Repositories,
) *Service {
	return &Service{
		ycUser:       &model.YCUser{},
		logger:       logger,
		config:       config,
		repositories: repositories,
	}
}
