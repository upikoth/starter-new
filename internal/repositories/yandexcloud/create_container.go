package yandexcloud

import (
	"context"
	"encoding/json"
	"net/http"

	"github.com/upikoth/starter-new/internal/model"
)

type createContainerRequest struct {
	FolderID string `json:"folderId"`
	Name     string `json:"name"`
}

type createContainerResponse struct {
	OperationID string `json:"id"`
	Done        bool   `json:"done"`
}

func (y *YandexCloud) CreateContainer(
	ctx context.Context,
	folderID string,
	containerName string,
) (*model.CreateContainerResponse, error) {
	y.logger.Info("Создаем container в yandex cloud")

	reqStruct := createContainerRequest{
		FolderID: folderID,
		Name:     containerName,
	}

	bodyBytes, err := y.sendHttpRequest(
		ctx,
		http.MethodPost,
		"https://serverless-containers.api.cloud.yandex.net/containers/v1/containers",
		reqStruct,
	)

	if err != nil {
		y.logger.Error(err.Error())
		return nil, err
	}

	resParsed := createContainerResponse{}
	err = json.Unmarshal(bodyBytes, &resParsed)

	if err != nil {
		y.logger.Error(err.Error())
		return nil, err
	}

	y.logger.Info("Container успешно создан!")

	return &model.CreateContainerResponse{
		OperationID: resParsed.OperationID,
		Done:        resParsed.Done,
	}, nil
}
