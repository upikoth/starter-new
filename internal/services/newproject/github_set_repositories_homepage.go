package newproject

import (
	"context"
	"github.com/upikoth/starter-new/internal/model"
	"golang.org/x/sync/errgroup"
)

func (p *Service) SetGithubRepositoriesHomepage(ctx context.Context) error {
	eg, newCtx := errgroup.WithContext(ctx)

	eg.Go(func() error {
		return p.repositories.Github.SetRepositoryHomepageURL(newCtx, model.SetGithubRepositoryHomepageURLRequest{
			GithubUserName: p.config.GitHub.UserName,
			GithubRepoName: p.newProject.GetBackendRepositoryName(),
			URL:            p.newProject.GetSwaggerDocsURL(),
		})
	})

	eg.Go(func() error {
		return p.repositories.Github.SetRepositoryHomepageURL(newCtx, model.SetGithubRepositoryHomepageURLRequest{
			GithubUserName: p.config.GitHub.UserName,
			GithubRepoName: p.newProject.GetFrontendRepositoryName(),
			URL:            p.newProject.GetDomainURL(),
		})
	})

	if err := eg.Wait(); err != nil {
		return err
	}

	p.logger.Info("Github: homepage url для frontend и backend проектов добавлены")

	return nil
}
