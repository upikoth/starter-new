package yandexcloud

import (
	"github.com/upikoth/starter-new/internal/config"
	"github.com/upikoth/starter-new/internal/constants"
	"github.com/upikoth/starter-new/internal/pkg/logger"
	"go.uber.org/ratelimit"
)

type YandexCloud struct {
	logger   logger.Logger
	config   *config.Config
	rl       ratelimit.Limiter
	iamToken string
}

func New(
	logger logger.Logger,
	config *config.Config,
) (*YandexCloud, error) {
	rl := ratelimit.New(constants.YandexCloudRPS)

	yc := &YandexCloud{
		logger: logger,
		config: config,
		rl:     rl,
	}

	err := yc.fillIamToken()

	if err != nil {
		return nil, err
	}

	return yc, nil
}
