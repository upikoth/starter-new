package newproject

import (
	"context"
	"errors"
)

func (p *Service) CreateYCLogGroup(ctx context.Context) error {
	res, err := p.repositories.YandexCloud.CreateLoggingGroup(
		ctx,
		p.newProject.folderID,
		p.getProjectLoggingGroupName(),
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

	p.newProject.loggingGroupID = res.LogGroupID

	return nil
}
