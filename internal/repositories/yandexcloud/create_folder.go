package yandexcloud

import (
	"context"
	"encoding/json"
)

type createFolderRequest struct {
	CloudID     string   `json:"cloudId"`
	Name        string   `json:"name"`
	Description string   `json:"description"`
	Labels      struct{} `json:"labels"`
}

type createFolderResponse struct {
	Response struct {
		FolderID string `json:"id"`
	} `json:"response"`
}

func (y *YandexCloud) CreateFolder(ctx context.Context, folderName string) (string, error) {
	y.logger.Info("Создаем folder в yandex cloud")

	reqStruct := createFolderRequest{
		CloudID:     y.config.YandexCloud.CloudID,
		Name:        folderName,
		Description: "",
	}

	bodyBytes, err := y.sendHttpPostRequest(
		context.Background(),
		"https://resource-manager.api.cloud.yandex.net/resource-manager/v1/folders",
		reqStruct,
	)

	if err != nil {
		y.logger.Error(err.Error())
		return "", err
	}

	resParsed := createFolderResponse{}
	err = json.Unmarshal(bodyBytes, &resParsed)

	if err != nil {
		y.logger.Error(err.Error())
		return "", err
	}

	y.logger.Info("Folder успешно создан!")

	return resParsed.Response.FolderID, nil
}
