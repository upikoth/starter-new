package yandexcloud

import (
	"context"
	"encoding/json"
	"net/http"

	"github.com/upikoth/starter-new/internal/model"
)

type createLoggingGroupRequest struct {
	FolderID        string `json:"folderId"`
	Name            string `json:"name"`
	RetentionPeriod string `json:"retentionPeriod"`
}

type createLoggingGroupResponse struct {
	OperationID string `json:"id"`
	Done        bool   `json:"done"`
	Metadata    struct {
		LogGroupID string `json:"logGroupId"`
	} `json:"metadata"`
}

func (y *YandexCloud) CreateLoggingGroup(ctx context.Context, folderID, logGroupName string) (*model.CreateLoggingGroupResponse, error) {
	y.logger.Info("Создаем logging group в yandex cloud")

	reqStruct := createLoggingGroupRequest{
		FolderID:        folderID,
		Name:            logGroupName,
		RetentionPeriod: "259200s", // 3 дня
	}

	bodyBytes, err := y.sendHttpRequest(
		ctx,
		http.MethodPost,
		"https://logging.api.cloud.yandex.net/logging/v1/logGroups",
		reqStruct,
	)

	if err != nil {
		y.logger.Error(err.Error())
		return nil, err
	}

	resParsed := createLoggingGroupResponse{}
	err = json.Unmarshal(bodyBytes, &resParsed)

	if err != nil {
		y.logger.Error(err.Error())
		return nil, err
	}

	y.logger.Info("Log group успешно создана!")

	return &model.CreateLoggingGroupResponse{
		OperationID: resParsed.OperationID,
		LogGroupID:  resParsed.Metadata.LogGroupID,
		Done:        resParsed.Done,
	}, nil
}
