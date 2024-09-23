package user

import (
	"github.com/upikoth/starter-new/internal/config"
	"github.com/upikoth/starter-new/internal/model"
	"github.com/upikoth/starter-new/internal/pkg/logger"
	"github.com/upikoth/starter-new/internal/repositories"
)

type User struct {
	YCUser       *model.YCUser
	logger       logger.Logger
	config       *config.Config
	repositories *repositories.Repositories
}

func New(
	logger logger.Logger,
	config *config.Config,
	repositories *repositories.Repositories,
) *User {
	return &User{
		logger:       logger,
		config:       config,
		repositories: repositories,
		YCUser:       &model.YCUser{},
	}
}
