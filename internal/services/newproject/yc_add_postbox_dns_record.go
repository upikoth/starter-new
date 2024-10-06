package newproject

import (
	"context"
	"errors"
	"fmt"
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
		model.GetPostboxVerificationRecordRequest{
			IdentityID:      p.newProject.GetYCPostboxAddressID(),
			YCUserCookie:    cookie,
			YCUserCSRFToken: csrfToken,
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
		isCreated = p.repositories.YandexCloud.GetOperationStatus(ctx, res.OperationID)
	}

	if !isCreated {
		err := errors.New("dns record в процессе создания, статус операции не завершен")
		return err
	}

	p.logger.Info("Postbox привязан к DNS")

	return nil
}
