package yandexcloudbrowser

import (
	"github.com/upikoth/starter-new/internal/config"
	"github.com/upikoth/starter-new/internal/pkg/logger"
)

type YandexCloudBrowser struct {
	logger logger.Logger
	config *config.Config
}

func New(
	logger logger.Logger,
	config *config.Config,
) *YandexCloudBrowser {
	ycb := &YandexCloudBrowser{
		logger: logger,
		config: config,
	}

	return ycb
}
