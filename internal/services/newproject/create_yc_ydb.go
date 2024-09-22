package newproject

import (
	"context"
	"errors"
)

func (p *NewProject) CreateYCYDB(ctx context.Context) error {
	p.logger.Info("Создаем ydb в yandex cloud")

	res, err := p.repositories.YandexCloud.CreateYDB(ctx, p.project.FolderID, p.project.GetProjectYDBName())

	if err != nil {
		p.logger.Error(err.Error())
		return err
	}

	isCreated := res.Done
	if !isCreated {
		isCreated = p.repositories.YandexCloud.GetOperationStatus(ctx, res.OperationID)
	}

	if !isCreated {
		err := errors.New("ydb в процессе создания, статус операции не завершен")
		p.logger.Error(err.Error())
		return err
	}

	p.project.DatabaseEndpoint = res.DatabaseEndpoint
	p.logger.Info("YDB в yandex cloud успешно создана")
	return nil
}
