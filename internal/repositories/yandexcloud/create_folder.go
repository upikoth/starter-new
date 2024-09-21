package yandexcloud

import (
	"context"
	"encoding/json"
	"net/http"
)

type createFolderRequest struct {
	CloudID     string   `json:"cloudId"`
	Name        string   `json:"name"`
	Description string   `json:"description"`
	Labels      struct{} `json:"labels"`
}

type createFolderResponse struct {
	OperationID string `json:"id"`
	Metadata    struct {
		FolderID string `json:"folderId"`
	} `json:"metadata"`
}

func (y *YandexCloud) CreateFolder(ctx context.Context, folderName string) (string, string, error) {
	y.logger.Info("Создаем folder в yandex cloud")

	reqStruct := createFolderRequest{
		CloudID:     y.config.YandexCloud.CloudID,
		Name:        folderName,
		Description: "",
	}

	bodyBytes, err := y.sendHttpRequest(
		ctx,
		http.MethodPost,
		"https://resource-manager.api.cloud.yandex.net/resource-manager/v1/folders",
		reqStruct,
	)

	if err != nil {
		y.logger.Error(err.Error())
		return "", "", err
	}

	resParsed := createFolderResponse{}
	err = json.Unmarshal(bodyBytes, &resParsed)

	if err != nil {
		y.logger.Error(err.Error())
		return "", "", err
	}

	y.logger.Info("Folder успешно создан!")

	return resParsed.OperationID, resParsed.Metadata.FolderID, nil
}
