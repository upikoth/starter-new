package yandexcloud

import (
	"context"
	"encoding/json"
	"net/http"

	"github.com/upikoth/starter-new/internal/model"
)

type createYDBRequest struct {
	FolderID           string                             `json:"folderId"`
	Name               string                             `json:"name"`
	ServerlessDatabase createYDBRequestServerlessDatabase `json:"serverlessDatabase"`
}

type createYDBRequestServerlessDatabase struct {
	StorageSizeLimit string `json:"storageSizeLimit"`
}

type createYDBResponse struct {
	OperationID string                    `json:"id"`
	Done        bool                      `json:"done"`
	Response    createYDBResponseResponse `json:"response"`
}

type createYDBResponseResponse struct {
	Endpoint string `json:"endpoint"`
}

func (y *YandexCloud) CreateYDB(
	ctx context.Context,
	folderID string,
	databaseName string,
) (*model.CreateYDBResponse, error) {
	reqStruct := createYDBRequest{
		FolderID: folderID,
		Name:     databaseName,
		ServerlessDatabase: createYDBRequestServerlessDatabase{
			StorageSizeLimit: "1073741824", // 1Gb
		},
	}

	bodyBytes, err := y.sendHttpRequest(
		ctx,
		http.MethodPost,
		"https://ydb.api.cloud.yandex.net/ydb/v1/databases",
		reqStruct,
	)

	if err != nil {
		return nil, err
	}

	resParsed := createYDBResponse{}
	err = json.Unmarshal(bodyBytes, &resParsed)

	if err != nil {
		return nil, err
	}

	return &model.CreateYDBResponse{
		OperationID:      resParsed.OperationID,
		DatabaseEndpoint: resParsed.Response.Endpoint,
		Done:             resParsed.Done,
	}, nil
}
