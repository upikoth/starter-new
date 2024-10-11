package sentry

import (
	"github.com/upikoth/starter-new/internal/config"
	"github.com/upikoth/starter-new/internal/constants"
	"github.com/upikoth/starter-new/internal/pkg/logger"
	"go.uber.org/ratelimit"
)

type Sentry struct {
	logger logger.Logger
	config *config.Config
	rl     ratelimit.Limiter
}

func New(
	logger logger.Logger,
	config *config.Config,
) *Sentry {
	rl := ratelimit.New(constants.SentryRPS)

	return &Sentry{
		logger: logger,
		config: config,
		rl:     rl,
	}
}
