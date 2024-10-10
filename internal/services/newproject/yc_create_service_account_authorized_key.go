package newproject

import (
	"context"
)

func (p *Service) YCCreateServiceAccountAuthorizedKey(ctx context.Context) error {
	key, err := p.repositories.YandexCloud.CreateServiceAccountAuthorizedKey(
		ctx,
		p.newProject.GetYCServiceAccountID(),
		"Ключ для авторизации в container registry, YDB",
	)

	if err != nil {
		return err
	}

	p.newProject.SetYCSAJSONCredentials(key)
	p.logger.Info("YC: ключ авторизации для container registry и YDB создан")

	return nil
}
