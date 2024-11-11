package yandexcloudbrowser

import (
	"github.com/upikoth/starter-new/internal/config"
	"github.com/upikoth/starter-new/internal/constants"
	"github.com/upikoth/starter-new/internal/pkg/logger"
	"go.uber.org/ratelimit"
)

type YandexCloudBrowser struct {
	logger logger.Logger
	config *config.Config
	rl     ratelimit.Limiter
}

func New(
	logger logger.Logger,
	config *config.Config,
) *YandexCloudBrowser {
	rl := ratelimit.New(constants.YandexCloudBrowserRPS, ratelimit.WithoutSlack)

	ycb := &YandexCloudBrowser{
		logger: logger,
		config: config,
		rl:     rl,
	}

	return ycb
}
