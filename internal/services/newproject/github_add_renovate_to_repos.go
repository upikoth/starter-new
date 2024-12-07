package newproject

import (
	"context"
	"golang.org/x/sync/errgroup"
)

func (p *Service) AddRenovateToGithubRepositories(ctx context.Context) error {
	eg, newCtx := errgroup.WithContext(ctx)

	eg.Go(func() error {
		return p.repositories.Github.AddRenovateToRepository(newCtx, p.newProject.GetGithubBackendRepositoryID())
	})

	eg.Go(func() error {
		return p.repositories.Github.AddRenovateToRepository(newCtx, p.newProject.GetGithubFrontendRepositoryID())
	})

	if err := eg.Wait(); err != nil {
		return err
	}

	p.logger.Info("Github: renovate для frontend и backend проектов добавлен")

	return nil
}
