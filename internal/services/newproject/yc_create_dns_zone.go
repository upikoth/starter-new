package newproject

import (
	"context"
	"github.com/pkg/errors"
)

func (p *Service) CreateYCDNSZone(ctx context.Context) error {
	res, err := p.repositories.YandexCloud.CreateDNSZone(
		ctx,
		p.newProject.GetYCFolderID(),
		p.newProject.GetYCDNSZoneName(),
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
		return errors.New("YC: dns zone в процессе создания, статус операции не завершен")
	}

	p.newProject.SetYCDNSZoneID(res.DNSZoneId)
	p.logger.Info("YC: DNS зона создана")

	return nil
}
