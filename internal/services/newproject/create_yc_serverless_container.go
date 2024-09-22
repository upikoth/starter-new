package newproject

import (
	"context"
	"errors"
)

func (p *NewProject) CreateYCServerlessContainer(ctx context.Context) error {
	p.logger.Info("Создаем новый serverless container")

	res, err := p.repositories.YandexCloud.CreateContainer(ctx, p.project.FolderID, p.project.GetProjectServerlessContainerName())

	if err != nil {
		p.logger.Error(err.Error())
		return err
	}

	isContainerCreated := res.Done
	if !isContainerCreated {
		isContainerCreated = p.repositories.YandexCloud.GetOperationStatus(ctx, res.OperationID)
	}

	if !isContainerCreated {
		err := errors.New("serverless container в процессе создания, статус операции не завершен")
		p.logger.Error(err.Error())
		return err
	}

	p.logger.Info("Serverless container в yandex cloud успешно создан")
	return nil
}
