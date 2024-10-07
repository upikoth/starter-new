package newproject

import (
	"context"
	"fmt"
	"github.com/upikoth/starter-new/internal/model"
)

func (p *Service) BindCertificateToDNSZone(ctx context.Context) error {
	cookie, err := p.ycUserService.GetYcUserCookie(ctx)

	if err != nil {
		return err
	}

	csrfToken, err := p.ycUserService.GetYcUserCSRFToken(ctx)

	if err != nil {
		return err
	}

	certificateChallenge, err := p.repositories.YandexCloudBrowser.GetCertificateChallenge(
		ctx,
		model.YCGetCertificateChallengeRequest{
			YCBrowserRequest: model.YCBrowserRequest{
				YCUserCookie:    cookie,
				YCUserCSRFToken: csrfToken,
			},
			CertificateID: p.newProject.GetYCCertificateID(),
		},
	)

	if err != nil {
		return err
	}

	req := model.YCBindCertificateToDNSRequest{
		YCBrowserRequest: model.YCBrowserRequest{
			YCUserCookie:    cookie,
			YCUserCSRFToken: csrfToken,
		},
		DNSZoneID:     p.newProject.GetYCDNSZoneID(),
		DNSRecordName: certificateChallenge.DNSName,
		DNSRecordText: certificateChallenge.DNSText,
		DNSRecordOwnerID: fmt.Sprintf(
			"%s:%s:%s",
			p.newProject.GetYCCertificateID(),
			certificateChallenge.ChallegeType,
			certificateChallenge.DNSName,
		),
	}

	err = p.repositories.YandexCloudBrowser.BindCertificateToDNS(ctx, req)

	if err != nil {
		return err
	}

	p.logger.Info("YC: для домена настроен let's encrypt сертификат")

	return nil
}
