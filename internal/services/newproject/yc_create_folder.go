package newproject

import (
	"context"
	"github.com/pkg/errors"
)

func (p *Service) CreateYCFolder(ctx context.Context) error {
	res, err := p.repositories.YandexCloud.CreateFolder(ctx, p.newProject.GetName())

	if err != nil {
		return err
	}

	isCreated := res.Done
	if !isCreated {
		isCreated = p.repositories.YandexCloud.GetOperationStatus(ctx, res.OperationID)
	}

	if !isCreated {
		return errors.New("YC: folder в процессе создания, статус операции не завершен")
	}

	p.newProject.SetYCFolderID(res.FolderId)
	p.logger.Info("YC: folder создан")

	return nil
}
