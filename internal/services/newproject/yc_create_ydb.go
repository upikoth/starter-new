package newproject

import (
	"context"
	"github.com/pkg/errors"
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
		isCreated, err = p.repositories.YandexCloud.GetOperationStatus(ctx, res.OperationID)

		if err != nil {
			return err
		}
	}

	if !isCreated {
		return errors.New("YC: YDB в процессе создания, статус операции не завершен")
	}

	p.newProject.SetYCYDBEndpoint(res.DatabaseEndpoint)
	p.logger.Info("YC: YDB создана")

	return nil
}
