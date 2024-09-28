package newproject

import (
	"context"
	"errors"
)

func (p *Service) CreateYCFolder(ctx context.Context) error {
	res, err := p.repositories.YandexCloud.CreateFolder(ctx, p.newProject.name)

	if err != nil {
		return err
	}

	isCreated := res.Done
	if !isCreated {
		isCreated = p.repositories.YandexCloud.GetOperationStatus(ctx, res.OperationID)
	}

	if !isCreated {
		err := errors.New("folder в процессе создания, статус операции не завершен")
		return err
	}

	p.newProject.folderID = res.FolderId

	return nil
}
