package repositories

import (
	"github.com/upikoth/starter-new/internal/config"
	"github.com/upikoth/starter-new/internal/pkg/logger"
	"github.com/upikoth/starter-new/internal/repositories/consoleinput"
	"github.com/upikoth/starter-new/internal/repositories/fileinput"
	"github.com/upikoth/starter-new/internal/repositories/github"
	"github.com/upikoth/starter-new/internal/repositories/yandexcloud"
	"github.com/upikoth/starter-new/internal/repositories/yandexcloudbrowser"
)

type Repositories struct {
	ConsoleInput       *consoleinput.ConsoleInput
	FileInput          *fileinput.FileInput
	Github             *github.Github
	YandexCloud        *yandexcloud.YandexCloud
	YandexCloudBrowser *yandexcloudbrowser.YandexCloudBrowser
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
		FileInput: fileinput.New(
			logger,
			config,
		),
		Github: github.New(
			logger,
			config,
		),
		YandexCloud: yandexCloudRepo,
		YandexCloudBrowser: yandexcloudbrowser.New(
			logger,
			config,
		),
	}, nil
}
