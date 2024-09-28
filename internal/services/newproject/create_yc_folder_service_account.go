package newproject

import (
	"context"
	"fmt"
)

func (p *Service) CreateYCFolderServiceAccount(ctx context.Context) error {
	accountName := fmt.Sprintf("%s-manager", p.newProject.name)

	serviceAccountID, err := p.repositories.YandexCloud.CreateServiceAccount(
		ctx,
		accountName,
		p.newProject.folderID,
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
		p.newProject.folderID,
		accountRoles,
	)

	if err != nil {
		return err
	}

	p.newProject.serviceAccountID = serviceAccountID

	return nil
}
