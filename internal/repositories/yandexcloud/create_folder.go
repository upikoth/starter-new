package yandexcloud

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"time"
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

	body, err := json.Marshal(reqStruct)
	if err != nil {
		y.logger.Error(err.Error())
		return "", err
	}

	newCtx, cancel := context.WithTimeout(ctx, time.Minute)
	defer cancel()

	req, err := http.NewRequestWithContext(newCtx, http.MethodPost, "https://resource-manager.api.cloud.yandex.net/resource-manager/v1/folders", bytes.NewBuffer(body))
	if err != nil {
		y.logger.Error(err.Error())
		return "", err
	}

	req.Header.Add("Authorization", fmt.Sprintf("Bearer %s", y.iamToken))

	client := &http.Client{}
	res, err := client.Do(req)
	if err != nil {
		y.logger.Error(err.Error())
		return "", err
	}

	defer res.Body.Close()

	if res.StatusCode != http.StatusOK {
		err := errors.New("не удалось создать folder")
		y.logger.Error(err.Error())
		return "", err
	}

	bodyBytes, err := io.ReadAll(res.Body)
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
