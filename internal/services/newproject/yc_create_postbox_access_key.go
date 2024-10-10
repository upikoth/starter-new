package newproject

import (
	"context"
)

func (p *Service) YCCreatePostboxAccessKey(ctx context.Context) error {
	res, err := p.repositories.YandexCloud.CreateAccessKey(
		ctx,
		p.newProject.GetYCServiceAccountID(),
		"Ключ для отправки писем через yandex postbox",
	)

	if err != nil {
		return err
	}

	p.newProject.SetYCPostboxUsername(res.AccessKeyID)
	p.newProject.SetYCPostboxPassword(res.AccessKeySecret)
	p.logger.Info("YC: access key для postbox создан")

	return nil
}
