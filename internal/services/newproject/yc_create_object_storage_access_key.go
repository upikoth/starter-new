package newproject

import (
	"context"
)

func (p *Service) YCCreateObjectStorageAccessKey(ctx context.Context) error {
	res, err := p.repositories.YandexCloud.CreateAccessKey(
		ctx,
		p.newProject.GetYCServiceAccountID(),
		"Ключ для загрузки vue приложения на s3",
	)

	if err != nil {
		return err
	}

	p.newProject.SetYCObjectStorageAccessKeyID(res.AccessKeyID)
	p.newProject.SetYCObjectStorageAccessKeySecret(res.AccessKeySecret)
	p.logger.Info("YC: access key для s3 создан")

	return nil
}
