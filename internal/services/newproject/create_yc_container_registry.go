package newproject

import (
	"context"
	"errors"
)

func (p *NewProject) CreateYCContainerRegistry(ctx context.Context) error {
	p.logger.Info("Создаем новый registry")

	res, err := p.repositories.YandexCloud.CreateRegistry(ctx, p.project.FolderID, p.project.GetProjectRegistryName())

	if err != nil {
		p.logger.Error(err.Error())
		return err
	}

	isCreated := res.Done
	if !isCreated {
		isCreated = p.repositories.YandexCloud.GetOperationStatus(ctx, res.OperationID)
	}

	if !isCreated {
		err := errors.New("registry в процессе создания, статус операции не завершен")
		p.logger.Error(err.Error())
		return err
	}

	p.project.RegistryID = res.RegistryID
	p.logger.Info("Registry в yandex cloud успешно создан")
	return nil
}
