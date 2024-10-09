package newproject

import (
	"context"
	"github.com/pkg/errors"
)

func (p *Service) CreateYCContainerRegistry(ctx context.Context) error {
	res, err := p.repositories.YandexCloud.CreateRegistry(
		ctx,
		p.newProject.GetYCFolderID(),
		p.newProject.GetYCContainerRegistryName(),
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
		return errors.New("YC: registry в процессе создания, статус операции не завершен")
	}

	p.newProject.SetYCContainerRegistryID(res.RegistryID)
	p.logger.Info("YC: Container registry создано")

	return nil
}
