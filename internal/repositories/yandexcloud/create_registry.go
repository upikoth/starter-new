package yandexcloud

import (
	"context"
	"encoding/json"
	"github.com/pkg/errors"
	"net/http"

	"github.com/upikoth/starter-new/internal/model"
)

type createRegistryRequest struct {
	FolderID string `json:"folderId"`
	Name     string `json:"name"`
}

type createRegistryResponse struct {
	OperationID string `json:"id"`
	Done        bool   `json:"done"`
	Metadata    struct {
		RegistryID string `json:"registryId"`
	} `json:"metadata"`
}

func (y *YandexCloud) CreateRegistry(
	ctx context.Context,
	folderID string,
	registryName string,
) (*model.CreateRegistryResponse, error) {
	reqStruct := createRegistryRequest{
		FolderID: folderID,
		Name:     registryName,
	}

	bodyBytes, err := y.sendHttpRequest(
		ctx,
		http.MethodPost,
		"https://container-registry.api.cloud.yandex.net/container-registry/v1/registries",
		reqStruct,
	)

	if err != nil {
		return nil, err
	}

	resParsed := createRegistryResponse{}
	err = json.Unmarshal(bodyBytes, &resParsed)

	if err != nil {
		return nil, errors.WithStack(err)
	}

	return &model.CreateRegistryResponse{
		YCResponse: model.YCResponse{
			OperationID: resParsed.OperationID,
			Done:        resParsed.Done,
		},
		RegistryID: resParsed.Metadata.RegistryID,
	}, nil
}
