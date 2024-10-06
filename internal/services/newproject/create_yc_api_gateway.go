package newproject

import (
	"context"
	"errors"
	"github.com/upikoth/starter-new/internal/model"
)

func (p *Service) CreateYCApiGateway(ctx context.Context) error {
	res, err := p.repositories.YandexCloud.CreateApiGateway(ctx, model.CreateApiGatewayRequest{
		FolderID:                 p.newProject.folderID,
		Name:                     p.getApiGatewayName(),
		LogGroupID:               p.newProject.loggingGroupID,
		ProjectCapitalizeName:    p.getCapitalizeName(),
		FrontendStaticBucketName: p.getObjectStorageFrontendStaticBucketName(),
		ServiceAccountID:         p.newProject.serviceAccountID,
		BackendContainerID:       p.newProject.backendContainerID,
	})

	if err != nil {
		return err
	}

	isCreated := res.Done
	if !isCreated {
		isCreated = p.repositories.YandexCloud.GetOperationStatus(ctx, res.OperationID)
	}

	if !isCreated {
		err := errors.New("api gateway в процессе создания, статус операции не завершен")
		return err
	}

	p.newProject.apiGatewayID = res.ApiGatewayID
	p.logger.Info("Api gateway создан")

	return nil
}
