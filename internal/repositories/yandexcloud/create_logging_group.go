package yandexcloud

import (
	"context"
	"encoding/json"
	"github.com/pkg/errors"
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
		return nil, err
	}

	resParsed := createLoggingGroupResponse{}
	err = json.Unmarshal(bodyBytes, &resParsed)

	if err != nil {
		return nil, errors.WithStack(err)
	}

	return &model.CreateLoggingGroupResponse{
		YCResponse: model.YCResponse{
			OperationID: resParsed.OperationID,
			Done:        resParsed.Done,
		},
		LogGroupID: resParsed.Metadata.LogGroupID,
	}, nil
}
