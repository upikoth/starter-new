package newproject

import (
	"context"
	"fmt"
	"github.com/pkg/errors"
	"github.com/upikoth/starter-new/internal/model"
	"strings"
)

func (p *Service) AddYCPostboxDNSRecord(ctx context.Context) error {
	cookie, err := p.ycUserService.GetYcUserCookie(ctx)

	if err != nil {
		return err
	}

	csrfToken, err := p.ycUserService.GetYcUserCSRFToken(ctx)

	if err != nil {
		return err
	}

	record, err := p.repositories.YandexCloudBrowser.GetPostboxVerificationRecord(
		ctx,
		model.YCGetPostboxVerificationRecordRequest{
			YCBrowserRequest: model.YCBrowserRequest{
				YCUserCookie:    cookie,
				YCUserCSRFToken: csrfToken,
			},
			IdentityID: p.newProject.GetYCPostboxAddressID(),
		},
	)

	if err != nil {
		return err
	}

	res, err := p.repositories.YandexCloud.AddDNSRecord(
		ctx,
		p.newProject.GetYCDNSZoneID(),
		model.DNSRecord{
			Type:  record.Type,
			Name:  strings.Replace(record.Name, fmt.Sprintf(".%s", p.newProject.GetDomain()), "", 1),
			Value: fmt.Sprintf(`"%s"`, record.Value),
		},
	)

	if err != nil {
		return err
	}

	isCreated := res.Done
	if !isCreated {
		isCreated, err = p.repositories.YandexCloud.GetOperationStatus(ctx, res.OperationID)

		if err != nil {
			return err
		}
	}

	if !isCreated {
		return errors.New("YC: dns record в процессе создания, статус операции не завершен")
	}

	p.logger.Info("YC: в DNS создана запись для верификации postbox")

	return nil
}
