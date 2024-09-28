package newproject

import (
	"context"
	"errors"

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

	return nil
}

func (p *Service) createYCStorageBucketWithSecrets(ctx context.Context) error {
	res, err := p.repositories.YandexCloud.CreateBucket(
		ctx,
		p.newProject.folderID,
		p.getObjectStorageSecretsBucketName(),
	)

	if err != nil {
		return err
	}

	isBucketCreated := res.Done
	if !isBucketCreated {
		isBucketCreated = p.repositories.YandexCloud.GetOperationStatus(ctx, res.OperationID)
	}

	if !isBucketCreated {
		err := errors.New("bucket в процессе создания, статус операции не завершен")
		return err
	}

	return nil
}

func (p *Service) createYCStorageBucketForFrontendStatic(ctx context.Context) error {
	res, err := p.repositories.YandexCloud.CreateBucket(
		ctx,
		p.newProject.folderID,
		p.getObjectStorageFrontendStaticBucketName(),
	)

	if err != nil {
		return err
	}

	isBucketCreated := res.Done
	if !isBucketCreated {
		isBucketCreated = p.repositories.YandexCloud.GetOperationStatus(ctx, res.OperationID)
	}

	if !isBucketCreated {
		err := errors.New("bucket в процессе создания, статус операции не завершен")
		return err
	}

	return nil
}
