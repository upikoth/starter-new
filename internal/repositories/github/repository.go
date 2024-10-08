package github

import (
	"github.com/upikoth/starter-new/internal/config"
	"github.com/upikoth/starter-new/internal/constants"
	"github.com/upikoth/starter-new/internal/pkg/logger"
	"go.uber.org/ratelimit"
)

type Github struct {
	logger logger.Logger
	config *config.Config
	rl     ratelimit.Limiter
}

func New(
	logger logger.Logger,
	config *config.Config,
) *Github {
	rl := ratelimit.New(constants.GithubRPS)

	return &Github{
		logger: logger,
		config: config,
		rl:     rl,
	}
}
