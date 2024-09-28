package newproject

import (
	"context"
	"errors"
)

func (p *Service) CreateYCDNSZone(ctx context.Context) error {
	res, err := p.repositories.YandexCloud.CreateDNSZone(
		ctx,
		p.newProject.folderID,
		p.getProjectDNSZoneName(),
	)

	if err != nil {
		return err
	}

	isCreated := res.Done
	if !isCreated {
		isCreated = p.repositories.YandexCloud.GetOperationStatus(ctx, res.OperationID)
	}

	if !isCreated {
		err := errors.New("dns zone в процессе создания, статус операции не завершен")
		return err
	}

	return nil
}
