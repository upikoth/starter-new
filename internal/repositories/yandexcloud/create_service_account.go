package yandexcloud

import (
	"context"
	"encoding/json"
	"github.com/pkg/errors"
	"net/http"
)

type createServiceAccountRequest struct {
	FolderID           string `json:"folderId"`
	AccountFolderID    string `json:"rolesFolderId"`
	AccountName        string `json:"name"`
	AccountDescription string `json:"description"`
}

type createServiceAccountResponse struct {
	Metadata struct {
		ServiceAccountID string `json:"serviceAccountId"`
	} `json:"metadata"`
}

func (y *YandexCloud) CreateServiceAccount(
	ctx context.Context,
	accountName string,
	folderID string,
) (string, error) {
	reqStruct := createServiceAccountRequest{
		FolderID:           folderID,
		AccountFolderID:    folderID,
		AccountName:        accountName,
		AccountDescription: "",
	}

	bodyBytes, err := y.sendHttpRequest(
		ctx,
		http.MethodPost,
		"https://iam.api.cloud.yandex.net/iam/v1/serviceAccounts",
		reqStruct,
	)

	if err != nil {
		return "", err
	}

	resParsed := createServiceAccountResponse{}
	err = json.Unmarshal(bodyBytes, &resParsed)

	if err != nil {
		return "", errors.WithStack(err)
	}

	return resParsed.Metadata.ServiceAccountID, nil
}
