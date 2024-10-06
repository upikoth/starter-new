package newproject

import (
	"context"
	"time"

	"github.com/upikoth/starter-new/internal/model"
)

func (p *Service) CreateYCCertificate(ctx context.Context) error {
	cookie, err := p.ycUserService.GetYcUserCookie(ctx)

	if err != nil {
		return err
	}

	csrfToken, err := p.ycUserService.GetYcUserCSRFToken(ctx)

	if err != nil {
		return err
	}

	req := model.CreateCertificateRequest{
		FolderID:        p.newProject.GetYCFolderID(),
		Domain:          p.newProject.GetDomain(),
		CertificateName: p.newProject.GetYCCertificateName(),
		YCUserCookie:    cookie,
		YCUserCSRFToken: csrfToken,
	}

	res, err := p.repositories.YandexCloudBrowser.CreateCertificate(ctx, req)

	if err != nil {
		return err
	}

	time.Sleep(time.Second * 5)

	p.newProject.SetYCCertificateID(res.CertificateID)
	p.logger.Info("Yandex cloud сертификат создан")

	return nil
}
