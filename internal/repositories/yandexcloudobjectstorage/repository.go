package yandexcloudobjectstorage

import (
	"github.com/upikoth/starter-new/internal/config"
	"github.com/upikoth/starter-new/internal/constants"
	"github.com/upikoth/starter-new/internal/pkg/logger"
	"go.uber.org/ratelimit"
)

type YandexCloudObjectStorage struct {
	logger logger.Logger
	config *config.Config
	rl     ratelimit.Limiter
}

func New(
	logger logger.Logger,
	config *config.Config,
) *YandexCloudObjectStorage {
	rl := ratelimit.New(constants.YandexCloudObjectStorageRPS)

	ycObjectStorage := &YandexCloudObjectStorage{
		logger: logger,
		config: config,
		rl:     rl,
	}

	return ycObjectStorage
}
