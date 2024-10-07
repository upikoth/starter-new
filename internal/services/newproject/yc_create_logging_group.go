package newproject

import (
	"context"
	"errors"
)

func (p *Service) CreateYCLogGroup(ctx context.Context) error {
	res, err := p.repositories.YandexCloud.CreateLoggingGroup(
		ctx,
		p.newProject.GetYCFolderID(),
		p.newProject.GetYCLoggingGroupName(),
	)

	if err != nil {
		return err
	}

	isCreated := res.Done
	if !isCreated {
		isCreated = p.repositories.YandexCloud.GetOperationStatus(ctx, res.OperationID)
	}

	if !isCreated {
		err := errors.New("лог группа в процессе создания, статус операции не завершен")
		return err
	}

	p.newProject.SetYCLoggingGroupID(res.LogGroupID)
	p.logger.Info("YC: лог группа создана")

	return nil
}
