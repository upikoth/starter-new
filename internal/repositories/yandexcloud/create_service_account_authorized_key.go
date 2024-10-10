package yandexcloud

import (
	"context"
	"encoding/json"
	"github.com/pkg/errors"
	"net/http"
	"time"
)

type createServiceAccountAuthorizedKeyRequest struct {
	ServiceAccountID string `json:"serviceAccountId"`
	Description      string `json:"description"`
	KeyAlgorithm     string `json:"keyAlgorithm"`
}

type createServiceAccountAuthorizedKeyResponse struct {
	Key struct {
		ID               string    `json:"id"`
		CreatedAt        time.Time `json:"createdAt"`
		Description      string    `json:"description"`
		KeyAlgorithm     string    `json:"keyAlgorithm"`
		PublicKey        string    `json:"publicKey"`
		ServiceAccountId string    `json:"serviceAccountId"`
	} `json:"key"`
	PrivateKey string `json:"privateKey"`
}

type keyJSONStructure struct {
	ID               string    `json:"id"`
	ServiceAccountId string    `json:"service_account_id"`
	CreatedAt        time.Time `json:"created_at"`
	KeyAlgorithm     string    `json:"key_algorithm"`
	PublicKey        string    `json:"public_key"`
	PrivateKey       string    `json:"private_key"`
}

func (y *YandexCloud) CreateServiceAccountAuthorizedKey(
	ctx context.Context,
	serviceAccountID string,
	description string,
) (string, error) {
	reqStruct := createServiceAccountAuthorizedKeyRequest{
		ServiceAccountID: serviceAccountID,
		Description:      description,
		KeyAlgorithm:     "RSA_4096",
	}

	bodyBytes, err := y.sendHttpRequest(
		ctx,
		http.MethodPost,
		"https://iam.api.cloud.yandex.net/iam/v1/keys",
		reqStruct,
	)

	if err != nil {
		return "", err
	}

	resParsed := createServiceAccountAuthorizedKeyResponse{}
	err = json.Unmarshal(bodyBytes, &resParsed)

	if err != nil {
		return "", errors.WithStack(err)
	}

	key := keyJSONStructure{
		ID:               resParsed.Key.ID,
		ServiceAccountId: resParsed.Key.ServiceAccountId,
		CreatedAt:        resParsed.Key.CreatedAt,
		KeyAlgorithm:     resParsed.Key.KeyAlgorithm,
		PublicKey:        resParsed.Key.PublicKey,
		PrivateKey:       resParsed.PrivateKey,
	}

	bytes, err := json.Marshal(key)

	if err != nil {
		return "", errors.WithStack(err)
	}

	return string(bytes), nil
}
