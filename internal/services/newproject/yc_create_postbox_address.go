package newproject

import (
	"context"
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"encoding/pem"
	"errors"
	"github.com/upikoth/starter-new/internal/model"
)

func (p *Service) CreateYCPostboxAddress(ctx context.Context) error {
	cookie, err := p.ycUserService.GetYcUserCookie(ctx)

	if err != nil {
		return err
	}

	csrfToken, err := p.ycUserService.GetYcUserCSRFToken(ctx)

	if err != nil {
		return err
	}

	bitSize := 2048

	key, err := rsa.GenerateKey(rand.Reader, bitSize)
	if err != nil {
		return err
	}

	keyPEM := pem.EncodeToMemory(
		&pem.Block{
			Type:  "RSA PRIVATE KEY",
			Bytes: x509.MarshalPKCS1PrivateKey(key),
		},
	)

	res, err := p.repositories.YandexCloudBrowser.CreatePostboxAddress(ctx, model.CreatePostboxAddressRequest{
		FolderID:        p.newProject.GetYCFolderID(),
		AddressName:     p.newProject.GetYCPostboxName(),
		YCUserCookie:    cookie,
		YCUserCSRFToken: csrfToken,
		PrivateKey:      string(keyPEM),
		Selector:        "mail",
		LogGroupID:      p.newProject.GetYCLoggingGroupID(),
	})

	if err != nil {
		return err
	}

	isCreated := res.Done
	if !isCreated {
		isCreated = p.repositories.YandexCloud.GetOperationStatus(ctx, res.OperationID)
	}

	if !isCreated {
		err := errors.New("postbox address в процессе создания, статус операции не завершен")
		return err
	}

	p.newProject.SetYCPostboxAddressID(res.PostboxAddressID)
	p.logger.Info("Postbox создан")

	return nil
}