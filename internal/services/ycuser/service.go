package ycuser

import (
	"github.com/upikoth/starter-new/internal/config"
	"github.com/upikoth/starter-new/internal/pkg/logger"
	"github.com/upikoth/starter-new/internal/repositories"
)

type Service struct {
	ycUser       *ycUser
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
		ycUser:       &ycUser{},
		logger:       logger,
		config:       config,
		repositories: repositories,
	}
}
