package yandexcloud

import (
	"context"
	"encoding/json"
	"github.com/pkg/errors"
	"net/http"
	"time"

	"github.com/upikoth/starter-new/internal/model"
)

type createAccessKeyRequest struct {
	ServiceAccountID string `json:"serviceAccountId"`
	Description      string `json:"description"`
}

type createAccessKeyResponse struct {
	AccessKey struct {
		ID               string    `json:"id"`
		ServiceAccountId string    `json:"serviceAccountId"`
		CreatedAt        time.Time `json:"createdAt"`
		Description      string    `json:"description"`
		KeyID            string    `json:"keyId"`
	} `json:"accessKey"`
	Secret string `json:"secret"`
}

func (y *YandexCloud) CreateAccessKey(
	ctx context.Context,
	serviceAccountID string,
	description string,
) (*model.CreateAccessKeyResponse, error) {
	reqStruct := createAccessKeyRequest{
		ServiceAccountID: serviceAccountID,
		Description:      description,
	}

	bodyBytes, err := y.sendHttpRequest(
		ctx,
		http.MethodPost,
		"https://iam.api.cloud.yandex.net/iam/aws-compatibility/v1/accessKeys",
		reqStruct,
	)

	if err != nil {
		return nil, err
	}

	resParsed := createAccessKeyResponse{}
	err = json.Unmarshal(bodyBytes, &resParsed)

	if err != nil {
		return nil, errors.WithStack(err)
	}

	return &model.CreateAccessKeyResponse{
		AccessKeyID:     resParsed.AccessKey.KeyID,
		AccessKeySecret: resParsed.Secret,
	}, nil
}
