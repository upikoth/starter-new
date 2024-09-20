package yandexcloud

import (
	"github.com/upikoth/starter-new/internal/config"
	"github.com/upikoth/starter-new/internal/pkg/logger"
)

type YandexCloud struct {
	logger   logger.Logger
	config   *config.Config
	iamToken string
}

func New(
	logger logger.Logger,
	config *config.Config,
) (*YandexCloud, error) {
	yc := &YandexCloud{
		logger: logger,
		config: config,
	}

	err := yc.fillIamToken()

	if err != nil {
		return nil, err
	}

	return yc, nil
}
