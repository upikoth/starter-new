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
		model.GetCertificateChallengeRequest{
			CertificateID:   p.newProject.certificateID,
			YCUserCookie:    cookie,
			YCUserCSRFToken: csrfToken,
		},
	)

	if err != nil {
		return err
	}

	req := model.BindCertificateToDNSRequest{
		DNSZoneID:        p.newProject.dnsZoneID,
		DNSRecordName:    certificateChallenge.DNSName,
		DNSRecordText:    certificateChallenge.DNSText,
		DNSRecordOwnerID: fmt.Sprintf("%s:%s:%s", p.newProject.certificateID, certificateChallenge.ChallegeType, certificateChallenge.DNSName),
		YCUserCookie:     cookie,
		YCUserCSRFToken:  csrfToken,
	}

	err = p.repositories.YandexCloudBrowser.BindCertificateToDNS(ctx, req)

	if err != nil {
		return err
	}

	return nil
}
