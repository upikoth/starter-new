package newproject

import (
	"context"
	"github.com/upikoth/starter-new/internal/model"
)

func (p *Service) CreateSentryBackendProject(ctx context.Context) error {
	err := p.repositories.Sentry.CreateProject(ctx, model.CreateSentryProjectRequest{
		ProjectName:     p.newProject.GetBackendRepositoryName(),
		ProjectPlatform: "go",
	})

	if err != nil {
		return err
	}

	p.logger.Info("Sentry: проект для backend создан")

	return nil
}
