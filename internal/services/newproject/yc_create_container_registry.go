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
		isCreated = p.repositories.YandexCloud.GetOperationStatus(ctx, res.OperationID)
	}

	if !isCreated {
		return errors.New("registry в процессе создания, статус операции не завершен")
	}

	p.newProject.SetYCContainerRegistryID(res.RegistryID)
	p.logger.Info("YC: Container registry создано")

	return nil
}
