package newproject

import (
	"context"
	"github.com/pkg/errors"
)

func (p *Service) UpdateYCAccessToRegistry(ctx context.Context) error {
	res, err := p.repositories.YandexCloud.UpdateServiceAccountAccessToRegistry(
		ctx,
		p.newProject.GetYCServiceAccountID(),
		p.newProject.GetYCContainerRegistryID(),
	)

	if err != nil {
		return err
	}

	isUpdated := res.Done
	if !isUpdated {
		isUpdated, err = p.repositories.YandexCloud.GetOperationStatus(ctx, res.OperationID)

		if err != nil {
			return err
		}
	}

	if !isUpdated {
		return errors.New("YC: права на доступ к container registry устанавливаются, статус операции не завершен")
	}

	p.logger.Info("YC: права на доступ к container registry установлены")

	return nil
}
