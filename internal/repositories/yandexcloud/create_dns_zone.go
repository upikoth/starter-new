package yandexcloud

import (
	"context"
	"encoding/json"
	"net/http"

	"github.com/upikoth/starter-new/internal/model"
)

type createDNSZoneRequest struct {
	FolderID         string   `json:"folderId"`
	Zone             string   `json:"zone"`
	PublicVisibility struct{} `json:"publicVisibility"`
}

type createDNSZoneResponse struct {
	OperationID string `json:"id"`
	Done        bool   `json:"done"`
}

func (y *YandexCloud) CreateDNSZone(
	ctx context.Context,
	folderID string,
	zoneName string,
) (*model.CreateDNSZoneResponse, error) {
	y.logger.Info("Создаем dns zone в yandex cloud")

	reqStruct := createDNSZoneRequest{
		FolderID:         folderID,
		Zone:             zoneName,
		PublicVisibility: struct{}{},
	}

	bodyBytes, err := y.sendHttpRequest(
		ctx,
		http.MethodPost,
		"https://dns.api.cloud.yandex.net/dns/v1/zones",
		reqStruct,
	)

	if err != nil {
		y.logger.Error(err.Error())
		return nil, err
	}

	resParsed := createDNSZoneResponse{}
	err = json.Unmarshal(bodyBytes, &resParsed)

	if err != nil {
		y.logger.Error(err.Error())
		return nil, err
	}

	y.logger.Info("Dns zone успешно создана!")

	return &model.CreateDNSZoneResponse{
		OperationID: resParsed.OperationID,
		Done:        resParsed.Done,
	}, nil
}
