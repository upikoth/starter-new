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
			p.newProject.GetYCCertificateID(),
		)

		if err != nil {
			return err
		}

		isCertificateNotIssued = certificate.Status != "ISSUED"

		p.logger.Info("YC: Сертификат не был выдан, ожидаем 1 минуту")
		retrays -= 1
		time.Sleep(time.Minute)
	}

	if isCertificateNotIssued {
		p.logger.Warn(fmt.Sprintf("YC: Сертификат не был выдан в течении %d минут, необходимо привязать api gateway к dns вручную", retrays))
		return nil
	}

	_, err = p.repositories.YandexCloud.AddDomainToGateway(
		ctx,
		p.newProject.GetDomain(),
		p.newProject.GetYCCertificateID(),
		p.newProject.GetYCAPIGatewayID(),
	)

	if err != nil {
		return err
	}

	apiGateway, err := p.repositories.YandexCloud.GetApiGateway(
		ctx,
		p.newProject.GetYCAPIGatewayID(),
	)

	if err != nil {
		return err
	}

	req := model.YCBindApiGatewayToDNSRequest{
		YCBrowserRequest: model.YCBrowserRequest{
			YCUserCookie:    cookie,
			YCUserCSRFToken: csrfToken,
		},
		DNSZoneID:        p.newProject.GetYCDNSZoneID(),
		DNSRecordName:    fmt.Sprintf("%s.", apiGateway.AttachedDomainName),
		DNSRecordText:    apiGateway.Domain,
		DNSRecordOwnerID: apiGateway.AttachedDomainID,
	}

	err = p.repositories.YandexCloudBrowser.BindApiGatewayToDNS(ctx, req)

	if err != nil {
		return err
	}

	p.logger.Info("YC: для домена настроен API gateway")

	return nil
}
