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

	req := model.YCCreateCertificateRequest{
		YCBrowserRequest: model.YCBrowserRequest{
			YCUserCookie:    cookie,
			YCUserCSRFToken: csrfToken,
		},
		FolderID:        p.newProject.GetYCFolderID(),
		Domain:          p.newProject.GetDomain(),
		CertificateName: p.newProject.GetYCCertificateName(),
	}

	res, err := p.repositories.YandexCloudBrowser.CreateCertificate(ctx, req)

	if err != nil {
		return err
	}

	time.Sleep(time.Second * 5)

	p.newProject.SetYCCertificateID(res.CertificateID)
	p.logger.Info("YC: сертификат let's encrypt создан")

	return nil
}
