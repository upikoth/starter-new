package newproject

import (
	"context"
	"errors"

	"golang.org/x/sync/errgroup"
)

func (p *NewProject) CreateYCStorageBuckets(ctx context.Context) error {
	p.logger.Info("Создаем object storage buckets в yandex cloud")

	eg, newCtx := errgroup.WithContext(ctx)

	eg.Go(func() error {
		return p.createYCStorageBucketWithSecrets(newCtx)
	})

	eg.Go(func() error {
		return p.createYCStorageBucketForFrontendStatic(newCtx)
	})

	err := eg.Wait()

	if err != nil {
		p.logger.Error("Не удалось создать object storage buckets в yandex cloud")
		return err
	}

	p.logger.Info("Object storage buckets в yandex cloud успешно создан")
	return nil
}

func (p *NewProject) createYCStorageBucketWithSecrets(ctx context.Context) error {
	p.logger.Info("Создаем object storage secrets bucket в yandex cloud")

	res, err := p.repositories.YandexCloud.CreateBucket(ctx, p.project.FolderID, p.project.GetObjectStorageSecretsBucketName())

	if err != nil {
		p.logger.Error(err.Error())
		return err
	}

	isBucketCreated := res.Done
	if !isBucketCreated {
		isBucketCreated = p.repositories.YandexCloud.GetOperationStatus(ctx, res.OperationID)
	}

	if !isBucketCreated {
		err := errors.New("bucket в процессе создания, статус операции не завершен")
		p.logger.Error(err.Error())
		return err
	}

	p.logger.Info("Object storage secrets bucket в yandex cloud успешно создан")
	return nil
}

func (p *NewProject) createYCStorageBucketForFrontendStatic(ctx context.Context) error {
	p.logger.Info("Создаем object storage static bucket в yandex cloud")

	res, err := p.repositories.YandexCloud.CreateBucket(
		ctx,
		p.project.FolderID,
		p.project.GetProjectSiteDomain(p.config.MainSiteDomainName),
	)

	if err != nil {
		p.logger.Error(err.Error())
		return err
	}

	isBucketCreated := res.Done
	if !isBucketCreated {
		isBucketCreated = p.repositories.YandexCloud.GetOperationStatus(ctx, res.OperationID)
	}

	if !isBucketCreated {
		err := errors.New("bucket в процессе создания, статус операции не завершен")
		p.logger.Error(err.Error())
		return err
	}

	p.logger.Info("Object storage static bucket в yandex cloud успешно создан")
	return nil
}
