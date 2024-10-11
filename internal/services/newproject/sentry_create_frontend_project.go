package newproject

import (
	"context"
	"github.com/upikoth/starter-new/internal/model"
)

func (p *Service) CreateSentryFrontendProject(ctx context.Context) error {
	err := p.repositories.Sentry.CreateProject(ctx, model.CreateSentryProjectRequest{
		ProjectName:     p.newProject.GetFrontendRepositoryName(),
		ProjectPlatform: "javascript-vue",
	})

	if err != nil {
		return err
	}

	p.logger.Info("Sentry: проект для frontend создан")

	return nil
}
