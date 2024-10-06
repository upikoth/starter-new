package newproject

import (
	"context"
	"errors"
)

func (p *Service) CreateYCServerlessContainer(ctx context.Context) error {
	res, err := p.repositories.YandexCloud.CreateContainer(
		ctx,
		p.newProject.folderID,
		p.getProjectServerlessContainerName(),
	)

	if err != nil {
		return err
	}

	isCreated := res.Done
	if !isCreated {
		isCreated = p.repositories.YandexCloud.GetOperationStatus(ctx, res.OperationID)
	}

	if !isCreated {
		err := errors.New("serverless container в процессе создания, статус операции не завершен")
		return err
	}

	p.newProject.backendContainerID = res.ContainerID
	p.logger.Info("Yandex cloud serverless container создан")

	return nil
}
