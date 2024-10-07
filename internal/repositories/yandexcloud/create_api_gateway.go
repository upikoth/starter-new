package yandexcloud

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/pkg/errors"
	"net/http"

	"github.com/upikoth/starter-new/internal/model"
)

type createApiGatewayRequest struct {
	FolderID    string                            `json:"folderId"`
	Name        string                            `json:"name"`
	LogOptions  createApiGatewayRequestLogOptions `json:"logOptions"`
	OpenapiSpec string                            `json:"openapiSpec"`
}

type createApiGatewayRequestLogOptions struct {
	LogGroupID string `json:"logGroupId"`
}

type createApiGatewayResponse struct {
	OperationID string `json:"id"`
	Done        bool   `json:"done"`
	Metadata    struct {
		ApiGatewayID string `json:"apiGatewayId"`
	} `json:"metadata"`
}

func (y *YandexCloud) CreateApiGateway(
	ctx context.Context,
	req model.YCCreateApiGatewayRequest,
) (*model.CreateApiGatewayResponse, error) {
	reqStruct := createApiGatewayRequest{
		FolderID: req.FolderID,
		Name:     req.Name,
		LogOptions: createApiGatewayRequestLogOptions{
			LogGroupID: req.LogGroupID,
		},
		OpenapiSpec: getOpeanapiSpec(req),
	}

	bodyBytes, err := y.sendHttpRequest(
		ctx,
		http.MethodPost,
		"https://serverless-apigateway.api.cloud.yandex.net/apigateways/v1/apigateways",
		reqStruct,
	)

	if err != nil {
		return nil, err
	}

	resParsed := createApiGatewayResponse{}
	err = json.Unmarshal(bodyBytes, &resParsed)

	if err != nil {
		return nil, errors.WithStack(err)
	}

	return &model.CreateApiGatewayResponse{
		YCResponse: model.YCResponse{
			OperationID: resParsed.OperationID,
			Done:        resParsed.Done,
		},
		ApiGatewayID: resParsed.Metadata.ApiGatewayID,
	}, nil
}

func getOpeanapiSpec(req model.YCCreateApiGatewayRequest) string {
	return fmt.Sprintf(`openapi: 3.0.0

# x-yc-apigateway:
#   smartWebSecurity:
#     securityProfileId: todo if needed

info:
  title: %s API
  version: 1.0.0

paths:
  /:
    get:
      x-yc-apigateway-integration:
        type: object_storage
        bucket: %s
        object: 'index.html'
        error_object: error.html
        service_account_id: %s

  /{file+}:
    get:
      summary: Serve static file from Yandex Cloud Object Storage
      parameters:
        - name: file
          in: path
          required: true
          schema:
            type: string
      x-yc-apigateway-integration:
        type: object_storage
        bucket: %s
        object: '{file}'
        error_object: error.html
        service_account_id: %s

  /api/{path+}:
    x-yc-apigateway-any-method:
      x-yc-apigateway-integration:
        type: serverless_containers
        service_account_id: %s
        container_id: %s
    parameters:
      - name: path
        in: path
        required: false
        schema:
          type: string`,
		req.ProjectCapitalizeName,
		req.FrontendStaticBucketName,
		req.ServiceAccountID,
		req.FrontendStaticBucketName,
		req.ServiceAccountID,
		req.ServiceAccountID,
		req.BackendContainerID,
	)
}
