package newproject

import (
	"context"
	"github.com/pkg/errors"
	"github.com/upikoth/starter-new/internal/model"
)

func (p *Service) CreateYCApiGateway(ctx context.Context) error {
	res, err := p.repositories.YandexCloud.CreateApiGateway(ctx, model.YCCreateApiGatewayRequest{
		FolderID:                 p.newProject.GetYCFolderID(),
		Name:                     p.newProject.GetYCApiGatewayName(),
		LogGroupID:               p.newProject.GetYCLoggingGroupID(),
		ProjectCapitalizeName:    p.newProject.GetCapitalizeName(),
		FrontendStaticBucketName: p.newProject.GetYCObjectStorageBucketNameStatic(),
		ServiceAccountID:         p.newProject.GetYCServiceAccountID(),
		BackendContainerID:       p.newProject.GetYCServerlessContainerID(),
	})

	if err != nil {
		return err
	}

	isCreated := res.Done
	if !isCreated {
		isCreated = p.repositories.YandexCloud.GetOperationStatus(ctx, res.OperationID)
	}

	if !isCreated {
		return errors.New("API gateway в процессе создания, статус операции не завершен")
	}

	p.newProject.SetYCAPIGatewayID(res.ApiGatewayID)
	p.logger.Info("YC: API gateway создан")

	return nil
}
