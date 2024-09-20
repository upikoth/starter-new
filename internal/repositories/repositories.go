package repositories

import (
	"github.com/upikoth/starter-new/internal/config"
	"github.com/upikoth/starter-new/internal/pkg/logger"
	"github.com/upikoth/starter-new/internal/repositories/consoleinput"
	"github.com/upikoth/starter-new/internal/repositories/github"
	"github.com/upikoth/starter-new/internal/repositories/yandexcloud"
)

type Repositories struct {
	ConsoleInput *consoleinput.ConsoleInput
	Github       *github.Github
	YandexCloud  *yandexcloud.YandexCloud
}

func New(
	logger logger.Logger,
	config *config.Config,
) (*Repositories, error) {
	yandexCloudRepo, err := yandexcloud.New(logger, config)

	if err != nil {
		return nil, err
	}

	return &Repositories{
		ConsoleInput: consoleinput.New(
			logger,
			config,
		),
		Github: github.New(
			logger,
			config,
		),
		YandexCloud: yandexCloudRepo,
	}, nil
}
