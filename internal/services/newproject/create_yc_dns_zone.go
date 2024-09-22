package newproject

import (
	"context"
	"errors"
)

func (p *NewProject) CreateYCDNSZone(ctx context.Context) error {
	p.logger.Info("Создаем dns zone")

	res, err := p.repositories.YandexCloud.CreateDNSZone(
		ctx,
		p.project.FolderID,
		p.project.GetProjectDNSZoneName(p.config.MainSiteDomainName),
	)

	if err != nil {
		p.logger.Error(err.Error())
		return err
	}

	isCreated := res.Done
	if !isCreated {
		isCreated = p.repositories.YandexCloud.GetOperationStatus(ctx, res.OperationID)
	}

	if !isCreated {
		err := errors.New("dns zone в процессе создания, статус операции не завершен")
		p.logger.Error(err.Error())
		return err
	}

	p.logger.Info("DNS zone в yandex cloud успешно создана")
	return nil
}
