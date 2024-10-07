package newproject

import (
	"context"
	"errors"
)

func (p *Service) CreateYCYDB(ctx context.Context) error {
	res, err := p.repositories.YandexCloud.CreateYDB(
		ctx,
		p.newProject.GetYCFolderID(),
		p.newProject.GetYCYDBName(),
	)

	if err != nil {
		return err
	}

	isCreated := res.Done
	if !isCreated {
		isCreated = p.repositories.YandexCloud.GetOperationStatus(ctx, res.OperationID)
	}

	if !isCreated {
		err := errors.New("YDB в процессе создания, статус операции не завершен")
		return err
	}

	p.newProject.SetYCYDBEndpoint(res.DatabaseEndpoint)
	p.logger.Info("YC: YDB создана")

	return nil
}
