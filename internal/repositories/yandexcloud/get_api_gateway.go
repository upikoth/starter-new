package yandexcloud

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/pkg/errors"
	"github.com/upikoth/starter-new/internal/model"
	"net/http"
)

type getApiGatewayResponse struct {
	AttachedDomains []struct {
		Enabled       bool   `json:"enabled"`
		DomainId      string `json:"domainId"`
		CertificateId string `json:"certificateId"`
		Domain        string `json:"domain"`
	} `json:"attachedDomains"`
	Domain string `json:"domain"`
}

func (y *YandexCloud) GetApiGateway(
	ctx context.Context,
	apiGatewayID string,
) (*model.ApiGateway, error) {
	bodyBytes, err := y.sendHttpRequest(
		ctx,
		http.MethodGet,
		fmt.Sprintf("https://serverless-apigateway.api.cloud.yandex.net/apigateways/v1/apigateways/%s", apiGatewayID),
		struct{}{},
	)

	if err != nil {
		return nil, err
	}

	resParsed := getApiGatewayResponse{}
	err = json.Unmarshal(bodyBytes, &resParsed)

	if err != nil {
		return nil, errors.WithStack(err)
	}

	return &model.ApiGateway{
		Domain:             resParsed.Domain,
		AttachedDomainID:   resParsed.AttachedDomains[0].DomainId,
		AttachedDomainName: resParsed.AttachedDomains[0].Domain,
	}, nil
}
