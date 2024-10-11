package newproject

import (
	"context"
	"github.com/upikoth/starter-new/internal/model"
)

func (p *Service) CreateSentryFrontendDSN(ctx context.Context) error {
	dsn, err := p.repositories.Sentry.CreateProjectDSN(ctx, model.CreateSentryProjectDSNRequest{
		ProjectName: p.newProject.GetFrontendRepositoryName(),
	})

	if err != nil {
		return err
	}

	p.logger.Info("Sentry: DSN для frontend создан")
	p.newProject.SetSentryFrontendDSN(dsn)

	return nil
}
