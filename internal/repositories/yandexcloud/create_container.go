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
	Metadata    struct {
		ContainerID string `json:"containerId"`
	} `json:"metadata"`
}

func (y *YandexCloud) CreateContainer(
	ctx context.Context,
	folderID string,
	containerName string,
) (*model.CreateContainerResponse, error) {
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
		return nil, err
	}

	resParsed := createContainerResponse{}
	err = json.Unmarshal(bodyBytes, &resParsed)

	if err != nil {
		return nil, err
	}

	return &model.CreateContainerResponse{
		YCResponse: model.YCResponse{
			OperationID: resParsed.OperationID,
			Done:        resParsed.Done,
		},
		ContainerID: resParsed.Metadata.ContainerID,
	}, nil
}
