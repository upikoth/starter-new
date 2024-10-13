package newproject

import (
	"context"
	"github.com/upikoth/starter-new/internal/model"
)

func (p *Service) UploadYCSecretsToObjectStorage(ctx context.Context) error {
	err := p.repositories.YandexCloudObjectStorage.UploadFileToBucket(
		ctx,
		model.UploadFileToBucketRequest{
			BucketName:  p.newProject.GetYCObjectStorageBucketNameSecrets(),
			FileName:    p.newProject.GetYCYDBFileName(),
			FileContent: []byte(p.newProject.GetYCSAJSONCredentials()),
			AccessKey:   p.newProject.GetYCObjectStorageAccessKeyID(),
			SecretKey:   p.newProject.GetYCObjectStorageAccessKeySecret(),
		},
	)

	if err != nil {
		return err
	}

	p.logger.Info("YC: файл с секретами в object storage успешно загружен")

	return nil
}
