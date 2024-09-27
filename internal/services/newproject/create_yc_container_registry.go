package newproject

import (
	"context"
	"errors"
)

func (p *NewProjectService) CreateYCContainerRegistry(ctx context.Context) error {
	res, err := p.repositories.YandexCloud.CreateRegistry(
		ctx,
		p.newProject.folderID,
		p.getProjectRegistryName(),
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

	p.newProject.registryID = res.RegistryID

	return nil
}
