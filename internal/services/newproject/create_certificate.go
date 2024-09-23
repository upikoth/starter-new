package newproject

import (
	"context"

	"github.com/upikoth/starter-new/internal/model"
)

func (p *NewProject) CreateYCCertificate(
	ctx context.Context,
	user *model.YCUser,
) error {
	p.logger.Info("Создаем certificate в yandex cloud")

	_, err := p.repositories.YandexCloudBrowser.CreateCertificate(
		ctx,
		p.project.FolderID,
		p.project.GetProjectSiteDomain(p.config.MainSiteDomainName),
		p.project.GetCertificateName(p.config.MainSiteDomainName),
		user,
	)

	if err != nil {
		p.logger.Error(err.Error())
		return err
	}

	p.logger.Info("Certificate в yandex cloud успешно создан")
	return nil
}
