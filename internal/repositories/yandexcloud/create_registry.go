package yandexcloud

import (
	"context"
	"encoding/json"
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
	y.logger.Info("Создаем registry в yandex cloud")

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
		y.logger.Error(err.Error())
		return nil, err
	}

	resParsed := createRegistryResponse{}
	err = json.Unmarshal(bodyBytes, &resParsed)

	if err != nil {
		y.logger.Error(err.Error())
		return nil, err
	}

	y.logger.Info("Registry успешно создан!")

	return &model.CreateRegistryResponse{
		OperationID: resParsed.OperationID,
		RegistryID:  resParsed.Metadata.RegistryID,
		Done:        resParsed.Done,
	}, nil
}
