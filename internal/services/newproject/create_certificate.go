package newproject

import (
	"context"

	"github.com/upikoth/starter-new/internal/model"
)

func (p *NewProjectService) CreateYCCertificate(ctx context.Context) error {
	cookie, err := p.ycUserService.GetYcUserCookie(ctx)

	if err != nil {
		return err
	}

	csrfToken, err := p.ycUserService.GetYcUserCSRFToken(ctx)

	if err != nil {
		return err
	}

	req := model.CreateCertificateRequest{
		FolderID:        p.newProject.folderID,
		Domain:          p.getProjectSiteDomain(),
		CertificateName: p.getCertificateName(),
		YCUserCookie:    cookie,
		YCUserCSRFToken: csrfToken,
	}

	_, err = p.repositories.YandexCloudBrowser.CreateCertificate(ctx, req)

	if err != nil {
		return err
	}

	return nil
}
