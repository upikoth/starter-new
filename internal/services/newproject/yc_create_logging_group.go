package newproject

import (
	"context"
	"github.com/pkg/errors"
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
		isCreated, err = p.repositories.YandexCloud.GetOperationStatus(ctx, res.OperationID)

		if err != nil {
			return err
		}
	}

	if !isCreated {
		return errors.New("YC: лог группа в процессе создания, статус операции не завершен")
	}

	p.newProject.SetYCLoggingGroupID(res.LogGroupID)
	p.logger.Info("YC: лог группа создана")

	return nil
}
