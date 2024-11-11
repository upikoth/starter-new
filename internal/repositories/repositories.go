package repositories

import (
	"github.com/upikoth/starter-new/internal/config"
	"github.com/upikoth/starter-new/internal/pkg/logger"
	"github.com/upikoth/starter-new/internal/repositories/consoleinput"
	"github.com/upikoth/starter-new/internal/repositories/fileinput"
	"github.com/upikoth/starter-new/internal/repositories/github"
	"github.com/upikoth/starter-new/internal/repositories/sentry"
	"github.com/upikoth/starter-new/internal/repositories/yandexcloud"
	"github.com/upikoth/starter-new/internal/repositories/yandexcloudbrowser"
	"github.com/upikoth/starter-new/internal/repositories/yandexcloudobjectstorage"
)

type Repositories struct {
	ConsoleInput             *consoleinput.ConsoleInput
	FileInput                *fileinput.FileInput
	Github                   *github.Github
	YandexCloud              *yandexcloud.YandexCloud
	YandexCloudBrowser       *yandexcloudbrowser.YandexCloudBrowser
	YandexCloudObjectStorage *yandexcloudobjectstorage.YandexCloudObjectStorage
	Sentry                   *sentry.Sentry
}

func New(
	log logger.Logger,
	cfg *config.Config,
) (*Repositories, error) {
	yandexCloudRepo, err := yandexcloud.New(log, cfg)

	if err != nil {
		return nil, err
	}

	return &Repositories{
		ConsoleInput: consoleinput.New(
			log,
			cfg,
		),
		FileInput: fileinput.New(
			log,
			cfg,
		),
		Github: github.New(
			log,
			cfg,
		),
		YandexCloud: yandexCloudRepo,
		YandexCloudBrowser: yandexcloudbrowser.New(
			log,
			cfg,
		),
		YandexCloudObjectStorage: yandexcloudobjectstorage.New(
			log,
			cfg,
		),
		Sentry: sentry.New(
			log,
			cfg,
		),
	}, nil
}
