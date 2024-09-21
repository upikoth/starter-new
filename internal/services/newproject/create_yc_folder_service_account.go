package newproject

import (
	"context"
	"fmt"
)

func (p *NewProject) CreateYCFolderServiceAccount(ctx context.Context) error {
	p.logger.Info("Создаем service account для folder в yandex cloud")

	accountName := fmt.Sprintf("%s-manager", p.project.Name)

	serviceAccountID, err := p.repositories.YandexCloud.CreateServiceAccount(
		ctx,
		accountName,
		p.project.FolderID,
	)

	if err != nil {
		p.logger.Error(err.Error())
		return err
	}

	accountRoles := []string{
		"serverless-containers.editor",
		"logging.editor",
		"iam.serviceAccounts.user",
		"ydb.editor",
		"postbox.sender",
		"storage.editor",
	}

	err = p.repositories.YandexCloud.UpdateServiceAccountRoles(
		ctx,
		serviceAccountID,
		p.project.FolderID,
		accountRoles,
	)

	if err != nil {
		p.logger.Error(err.Error())
		return err
	}

	p.project.ServiceAccountID = serviceAccountID
	p.logger.Info("Service account в yandex cloud успешно создан, роли заданы")
	return nil
}
