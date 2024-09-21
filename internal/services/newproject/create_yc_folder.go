package newproject

import (
	"context"
	"errors"
	"time"
)

func (p *NewProject) CreateYCFolder(ctx context.Context) error {
	p.logger.Info("Создаем новую папку проекта в yandex cloud")

	operationID, folderId, err := p.repositories.YandexCloud.CreateFolder(ctx, p.project.Name)

	if err != nil {
		p.logger.Error(err.Error())
		return err
	}

	isFolderCreated := false
	for i := 0; i < 10; i += 1 {
		time.Sleep(time.Second)
		done, err := p.repositories.YandexCloud.GetOperationStatus(ctx, operationID)

		if err != nil {
			p.logger.Error(err.Error())
			return err
		}

		if done {
			isFolderCreated = true
			break
		}
	}

	if !isFolderCreated {
		err := errors.New("folder в процессе создания, не удалось создать за 10 секунд")
		p.logger.Error(err.Error())
		return err
	}

	p.project.FolderID = folderId
	p.logger.Info("Folder в yandex cloud успешно создан")
	return nil
}
