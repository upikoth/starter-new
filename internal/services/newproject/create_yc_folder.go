package newproject

import (
	"context"
	"errors"
)

func (p *NewProject) CreateYCFolder(ctx context.Context) error {
	p.logger.Info("Создаем новую папку проекта в yandex cloud")

	res, err := p.repositories.YandexCloud.CreateFolder(ctx, p.project.Name)

	if err != nil {
		p.logger.Error(err.Error())
		return err
	}

	isCreated := res.Done
	if !isCreated {
		isCreated = p.repositories.YandexCloud.GetOperationStatus(ctx, res.OperationID)
	}

	if !isCreated {
		err := errors.New("folder в процессе создания, статус операции не завершен")
		p.logger.Error(err.Error())
		return err
	}

	p.project.FolderID = res.FolderId
	p.logger.Info("Folder в yandex cloud успешно создан")
	return nil
}
