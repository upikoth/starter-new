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
		return p.createLocalBackendRepo(newCtx)
	})

	eg.Go(func() error {
		return p.createLocalFrontendRepo(newCtx)
	})

	if err := eg.Wait(); err != nil {
		return err
	}

	p.logger.Info("Local: локальные папки для frontend и backend созданы")

	return nil
}

func (p *Service) createLocalProjectDirectories(ctx context.Context) error {
	eg, _ := errgroup.WithContext(ctx)

	if err := os.Mkdir(p.newProject.GetLocalPath(), 0777); err != nil {
		return err
	}

	eg.Go(func() error {
		return os.Mkdir(p.newProject.GetBackendLocalPath(), 0777)
	})

	eg.Go(func() error {
		return os.Mkdir(p.newProject.GetFrontendLocalPath(), 0777)
	})

	if err := eg.Wait(); err != nil {
		return err
	}

	return nil
}
