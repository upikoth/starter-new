package newproject

import (
	"context"
	"errors"
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
		isCreated = p.repositories.YandexCloud.GetOperationStatus(ctx, res.OperationID)
	}

	if !isCreated {
		err := errors.New("registry в процессе создания, статус операции не завершен")
		return err
	}

	p.newProject.SetYCContainerRegistryID(res.RegistryID)
	p.logger.Info("YC: Container registry создано")

	return nil
}
