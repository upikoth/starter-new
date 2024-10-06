package newproject

import (
	"context"

	"golang.org/x/sync/errgroup"
)

func (p *Service) CreateGithubRepositories(ctx context.Context) error {
	eg, newCtx := errgroup.WithContext(ctx)

	eg.Go(func() error {
		return p.repositories.Github.CreateRepository(newCtx, p.getBackendRepoName())
	})

	eg.Go(func() error {
		return p.repositories.Github.CreateRepository(newCtx, p.getFrontendRepoName())
	})

	if err := eg.Wait(); err != nil {
		return err
	}

	p.logger.Info("Репозитории в github для frontend и backend проектов созданы")

	return nil
}
