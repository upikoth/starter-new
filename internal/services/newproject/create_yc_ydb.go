package newproject

import (
	"context"
	"errors"
)

func (p *Service) CreateYCYDB(ctx context.Context) error {
	res, err := p.repositories.YandexCloud.CreateYDB(
		ctx,
		p.newProject.folderID,
		p.getProjectYDBName(),
	)

	if err != nil {
		return err
	}

	isCreated := res.Done
	if !isCreated {
		isCreated = p.repositories.YandexCloud.GetOperationStatus(ctx, res.OperationID)
	}

	if !isCreated {
		err := errors.New("ydb в процессе создания, статус операции не завершен")
		return err
	}

	p.newProject.databaseEndpoint = res.DatabaseEndpoint

	return nil
}
