package newproject

import (
	"context"
	"fmt"
	"github.com/upikoth/starter-new/internal/model"
	"time"
)

func (p *Service) BindYCGatewayToDNS(ctx context.Context) error {
	cookie, err := p.ycUserService.GetYcUserCookie(ctx)

	if err != nil {
		return err
	}

	csrfToken, err := p.ycUserService.GetYcUserCSRFToken(ctx)

	if err != nil {
		return err
	}

	isCertificateNotIssued := true
	retrays := 20

	for isCertificateNotIssued && retrays > 0 {
		certificate, err := p.repositories.YandexCloud.GetCertificate(
			ctx,
			p.newProject.certificateID,
		)

		if err != nil {
			return err
		}

		isCertificateNotIssued = certificate.Status != "ISSUED"

		p.logger.Info("Сертификат не был выдан, ожидаем 1 минуту")
		retrays -= 1
		time.Sleep(time.Minute)
	}

	if isCertificateNotIssued {
		p.logger.Warn(fmt.Sprintf("Сертификат не был выдан в течении %s минут, необходимо привязать api gateway к dns вручную", retrays))
		return nil
	}

	_, err = p.repositories.YandexCloud.AddDomainToGateway(
		ctx,
		p.getProjectSiteDomain(),
		p.newProject.certificateID,
		p.newProject.apiGatewayID,
	)

	if err != nil {
		return err
	}

	apiGateway, err := p.repositories.YandexCloud.GetApiGateway(
		ctx,
		p.newProject.apiGatewayID,
	)

	if err != nil {
		return err
	}

	req := model.BindApiGatewayToDNSRequest{
		DNSZoneID:        p.newProject.dnsZoneID,
		DNSRecordName:    fmt.Sprintf("%s.", apiGateway.AttachedDomainName),
		DNSRecordText:    apiGateway.Domain,
		DNSRecordOwnerID: apiGateway.AttachedDomainID,
		YCUserCookie:     cookie,
		YCUserCSRFToken:  csrfToken,
	}

	err = p.repositories.YandexCloudBrowser.BindApiGatewayToDNS(ctx, req)

	if err != nil {
		return err
	}

	return nil
}
