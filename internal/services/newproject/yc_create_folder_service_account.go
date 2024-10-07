package newproject

import (
	"context"
	"fmt"
)

func (p *Service) CreateYCFolderServiceAccount(ctx context.Context) error {
	accountName := fmt.Sprintf("%s-manager", p.newProject.GetName())

	serviceAccountID, err := p.repositories.YandexCloud.CreateServiceAccount(
		ctx,
		accountName,
		p.newProject.GetYCFolderID(),
	)

	if err != nil {
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
		p.newProject.GetYCFolderID(),
		accountRoles,
	)

	if err != nil {
		return err
	}

	p.newProject.SetYCServiceAccountID(serviceAccountID)
	p.logger.Info("YC: Сервисный аккаунт в для нового проекта создан")

	return nil
}
