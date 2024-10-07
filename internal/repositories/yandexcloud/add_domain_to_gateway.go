package yandexcloud

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/upikoth/starter-new/internal/model"
)

type addDomainToGatewayRequest struct {
	DomainName    string `json:"domainName"`
	CertificateID string `json:"certificateId"`
}

type addDomainToGatewayResponse struct {
	OperationID string `json:"id"`
	Done        bool   `json:"done"`
	Metadata    struct {
		ContainerID string `json:"containerId"`
	} `json:"metadata"`
}

func (y *YandexCloud) AddDomainToGateway(
	ctx context.Context,
	domainName string,
	certificateID string,
	apiGatewayID string,
) (*model.AddDomainToGatewayResponse, error) {
	reqStruct := addDomainToGatewayRequest{
		DomainName:    domainName,
		CertificateID: certificateID,
	}

	bodyBytes, err := y.sendHttpRequest(
		ctx,
		http.MethodPost,
		fmt.Sprintf("https://serverless-apigateway.api.cloud.yandex.net/apigateways/v1/apigateways/%s:addDomain", apiGatewayID),
		reqStruct,
	)

	if err != nil {
		return nil, err
	}

	resParsed := addDomainToGatewayResponse{}
	err = json.Unmarshal(bodyBytes, &resParsed)

	if err != nil {
		return nil, err
	}

	return &model.AddDomainToGatewayResponse{
		YCResponse: model.YCResponse{
			OperationID: resParsed.OperationID,
			Done:        resParsed.Done,
		},
	}, nil
}
