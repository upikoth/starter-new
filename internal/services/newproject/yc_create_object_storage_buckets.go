package newproject

import (
	"context"
	"github.com/pkg/errors"

	"golang.org/x/sync/errgroup"
)

func (p *Service) CreateYCStorageBuckets(ctx context.Context) error {
	eg, newCtx := errgroup.WithContext(ctx)

	eg.Go(func() error {
		return p.createYCStorageBucketWithSecrets(newCtx)
	})

	eg.Go(func() error {
		return p.createYCStorageBucketForFrontendStatic(newCtx)
	})

	err := eg.Wait()

	if err != nil {
		return err
	}

	p.logger.Info("YC: s3 бакеты для билда фронта и хранения секретов созданы")

	return nil
}

func (p *Service) createYCStorageBucketWithSecrets(ctx context.Context) error {
	res, err := p.repositories.YandexCloud.CreateBucket(
		ctx,
		p.newProject.GetYCFolderID(),
		p.newProject.GetYCObjectStorageBucketNameSecrets(),
	)

	if err != nil {
		return err
	}

	isBucketCreated := res.Done
	if !isBucketCreated {
		isBucketCreated, err = p.repositories.YandexCloud.GetOperationStatus(ctx, res.OperationID)

		if err != nil {
			return err
		}
	}

	if !isBucketCreated {
		return errors.New("bucket в процессе создания, статус операции не завершен")
	}

	return nil
}

func (p *Service) createYCStorageBucketForFrontendStatic(ctx context.Context) error {
	res, err := p.repositories.YandexCloud.CreateBucket(
		ctx,
		p.newProject.GetYCFolderID(),
		p.newProject.GetYCObjectStorageBucketNameStatic(),
	)

	if err != nil {
		return err
	}

	isBucketCreated := res.Done
	if !isBucketCreated {
		isBucketCreated, err = p.repositories.YandexCloud.GetOperationStatus(ctx, res.OperationID)

		if err != nil {
			return err
		}
	}

	if !isBucketCreated {
		return errors.New("YC: bucket в процессе создания, статус операции не завершен")
	}

	return nil
}
