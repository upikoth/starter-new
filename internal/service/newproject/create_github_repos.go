package newproject

import (
	"context"

	"golang.org/x/sync/errgroup"
)

func (p *NewProject) CreateGithubRepositories(ctx context.Context) error {
	p.logger.Info("Создаем репозитории в github")

	eg, newCtx := errgroup.WithContext(ctx)

	eg.Go(func() error {
		return p.repository.Github.CreateRepository(newCtx, p.project.GetBackendRepoName())
	})

	eg.Go(func() error {
		return p.repository.Github.CreateRepository(newCtx, p.project.GetFrontendRepoName())
	})

	if err := eg.Wait(); err != nil {
		p.logger.Error(err.Error())
		return err
	}

	p.logger.Info("Репозитории в github успешно созданы")
	return nil
}
