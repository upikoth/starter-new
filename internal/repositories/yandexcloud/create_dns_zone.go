package yandexcloud

import (
	"context"
	"encoding/json"
	"github.com/pkg/errors"
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
	Response    struct {
		ZoneID string `json:"id"`
	} `json:"response"`
}

func (y *YandexCloud) CreateDNSZone(
	ctx context.Context,
	folderID string,
	zoneName string,
) (*model.CreateDNSZoneResponse, error) {
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
		return nil, err
	}

	resParsed := createDNSZoneResponse{}
	err = json.Unmarshal(bodyBytes, &resParsed)

	if err != nil {
		return nil, errors.WithStack(err)
	}

	return &model.CreateDNSZoneResponse{
		YCResponse: model.YCResponse{
			OperationID: resParsed.OperationID,
			Done:        resParsed.Done,
		},
		DNSZoneId: resParsed.Response.ZoneID,
	}, nil
}
