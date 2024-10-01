package yandexcloud

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/upikoth/starter-new/internal/model"
)

type getCertificateResponse struct {
	Status string `json:"status"`
}

func (y *YandexCloud) GetCertificate(
	ctx context.Context,
	certificateID string,
) (*model.Certificate, error) {
	bodyBytes, err := y.sendHttpRequest(
		ctx,
		http.MethodGet,
		fmt.Sprintf("https://certificate-manager.api.cloud.yandex.net/certificate-manager/v1/certificates/%s", certificateID),
		struct{}{},
	)

	if err != nil {
		return nil, err
	}

	resParsed := getCertificateResponse{}
	err = json.Unmarshal(bodyBytes, &resParsed)

	if err != nil {
		return nil, err
	}

	return &model.Certificate{
		Status: resParsed.Status,
	}, nil
}
