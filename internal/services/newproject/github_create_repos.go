package newproject

import (
	"context"
	"golang.org/x/sync/errgroup"
)

func (p *Service) CreateGithubRepositories(ctx context.Context) error {
	eg, newCtx := errgroup.WithContext(ctx)

	var backendRepoID int
	var frontendRepoID int

	eg.Go(func() error {
		repoID, err := p.repositories.Github.CreateRepository(newCtx, p.newProject.GetBackendRepositoryName())

		if err != nil {
			return err
		}

		backendRepoID = repoID
		return nil
	})

	eg.Go(func() error {
		repoID, err := p.repositories.Github.CreateRepository(newCtx, p.newProject.GetFrontendRepositoryName())

		if err != nil {
			return err
		}

		frontendRepoID = repoID
		return nil
	})

	if err := eg.Wait(); err != nil {
		return err
	}

	p.newProject.SetGithubBackendRepositoryID(backendRepoID)
	p.newProject.SetGithubFrontendRepositoryID(frontendRepoID)

	p.logger.Info("Github: репозитории для frontend и backend проектов созданы")

	return nil
}
