package yandexcloud

import (
	"context"
	"encoding/json"
	"github.com/pkg/errors"
	"net/http"

	"github.com/upikoth/starter-new/internal/model"
)

type createBucketRequest struct {
	FolderID    string `json:"folderId"`
	Name        string `json:"name"`
	MaxSizeBits string `json:"maxSize"`
}

type createBucketResponse struct {
	OperationID string `json:"id"`
	Done        bool   `json:"done"`
}

func (y *YandexCloud) CreateBucket(
	ctx context.Context,
	folderID string,
	bucketName string,
) (*model.CreateBucketResponse, error) {
	reqStruct := createBucketRequest{
		FolderID:    folderID,
		Name:        bucketName,
		MaxSizeBits: "52428800", // 50Mb
	}

	bodyBytes, err := y.sendHttpRequest(
		ctx,
		http.MethodPost,
		"https://storage.api.cloud.yandex.net/storage/v1/buckets",
		reqStruct,
	)

	if err != nil {
		return nil, err
	}

	resParsed := createBucketResponse{}
	err = json.Unmarshal(bodyBytes, &resParsed)

	if err != nil {
		return nil, errors.WithStack(err)
	}

	return &model.CreateBucketResponse{
		YCResponse: model.YCResponse{
			OperationID: resParsed.OperationID,
			Done:        resParsed.Done,
		},
	}, nil
}
