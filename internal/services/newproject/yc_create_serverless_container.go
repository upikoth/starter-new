package newproject

import (
	"context"
	"github.com/pkg/errors"
)

func (p *Service) CreateYCServerlessContainer(ctx context.Context) error {
	res, err := p.repositories.YandexCloud.CreateContainer(
		ctx,
		p.newProject.GetYCFolderID(),
		p.newProject.GetYCServerlessContainerName(),
	)

	if err != nil {
		return err
	}

	isCreated := res.Done
	if !isCreated {
		isCreated, err = p.repositories.YandexCloud.GetOperationStatus(ctx, res.OperationID)

		if err != nil {
			return err
		}
	}

	if !isCreated {
		return errors.New("YC: serverless container в процессе создания, статус операции не завершен")
	}

	p.newProject.SetYCServerlessContainerID(res.ContainerID)
	p.logger.Info("YC: serverless container создан")

	return nil
}
