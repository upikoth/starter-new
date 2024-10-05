package newproject

import (
	"context"
	"golang.org/x/sync/errgroup"
	"os"
)

func (p *Service) CreateLocalRepos(ctx context.Context) error {
	err := p.createLocalProjectDirectories(ctx)

	if err != nil {
		return err
	}

	eg, newCtx := errgroup.WithContext(ctx)

	eg.Go(func() error {
		return p.createBackendRepo(newCtx)
	})

	eg.Go(func() error {
		return p.createFrontendRepo(newCtx)
	})

	if err := eg.Wait(); err != nil {
		return err
	}

	return nil
}

func (p *Service) createLocalProjectDirectories(_ context.Context) error {
	err := os.Mkdir(p.getProjectLocalPath(), 0777)

	if err != nil {
		return err
	}

	err = os.Mkdir(p.getProjectLocalPathBackend(), 0777)

	if err != nil {
		return err
	}

	err = os.Mkdir(p.getProjectLocalPathFrontend(), 0777)

	if err != nil {
		return err
	}

	return nil
}
