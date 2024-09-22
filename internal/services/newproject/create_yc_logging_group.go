package newproject

import (
	"context"
	"errors"
)

func (p *NewProject) CreateYCLogGroup(ctx context.Context) error {
	p.logger.Info("Создаем новую лог группу")

	res, err := p.repositories.YandexCloud.CreateLoggingGroup(ctx, p.project.FolderID, p.project.GetProjectLoggingGroupName())

	if err != nil {
		p.logger.Error(err.Error())
		return err
	}

	isCreated := res.Done
	if !isCreated {
		isCreated = p.repositories.YandexCloud.GetOperationStatus(ctx, res.OperationID)
	}

	if !isCreated {
		err := errors.New("лог группа в процессе создания, статус операции не завершен")
		p.logger.Error(err.Error())
		return err
	}

	p.project.LoggingGroupID = res.LogGroupID
	p.logger.Info("Лог группа в yandex cloud успешно создан")
	return nil
}
