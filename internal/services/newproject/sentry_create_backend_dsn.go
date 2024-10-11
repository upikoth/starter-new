package newproject

import (
	"context"
	"github.com/upikoth/starter-new/internal/model"
)

func (p *Service) CreateSentryBackendDSN(ctx context.Context) error {
	dsn, err := p.repositories.Sentry.CreateProjectDSN(ctx, model.CreateSentryProjectDSNRequest{
		ProjectName: p.newProject.GetBackendRepositoryName(),
	})

	if err != nil {
		return err
	}

	p.logger.Info("Sentry: DSN для backend создан")
	p.newProject.SetSentryBackendDSN(dsn)

	return nil
}
