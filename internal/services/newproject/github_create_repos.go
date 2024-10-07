package newproject

import (
	"context"
	"golang.org/x/sync/errgroup"
)

func (p *Service) CreateGithubRepositories(ctx context.Context) error {
	eg, newCtx := errgroup.WithContext(ctx)

	eg.Go(func() error {
		return p.repositories.Github.CreateRepository(newCtx, p.newProject.GetBackendRepositoryName())
	})

	eg.Go(func() error {
		return p.repositories.Github.CreateRepository(newCtx, p.newProject.GetFrontendRepositoryName())
	})

	if err := eg.Wait(); err != nil {
		return err
	}

	p.logger.Info("Github: репозитории для frontend и backend проектов созданы")

	return nil
}
