package yandexcloudbrowser

import (
	"context"
	"encoding/json"
	"github.com/pkg/errors"
	"net/http"

	"github.com/upikoth/starter-new/internal/model"
)

type createPostboxAddressRequest struct {
	FolderID          string                                       `json:"folderId"`
	Address           string                                       `json:"address"`
	SigningAttributes createPostboxAddressRequestSigningAttributes `json:"signingAttributes"`
	LogOptions        createPostboxAddressRequestLogOptions        `json:"logOptions"`
}

type createPostboxAddressRequestSigningAttributes struct {
	External createPostboxAddressRequestSigningAttributesExternal `json:"external"`
}

type createPostboxAddressRequestSigningAttributesExternal struct {
	Selector   string `json:"selector"`
	PrivateKey string `json:"privateKey"`
}

type createPostboxAddressRequestLogOptions struct {
	LogGroupID          string `json:"logGroupId"`
	IncludeMailStatuses bool   `json:"includeMailStatuses"`
}

type createPostboxAddressResponse struct {
	OperationID string `json:"id"`
	Done        bool   `json:"done"`
	Response    struct {
		ID string `json:"id"`
	} `json:"response"`
}

func (y *YandexCloudBrowser) CreatePostboxAddress(
	ctx context.Context,
	req model.YCCreatePostboxAddressRequest,
) (*model.CreatePostboxAddressResponse, error) {
	reqStruct := createPostboxAddressRequest{
		FolderID: req.FolderID,
		Address:  req.AddressName,
		SigningAttributes: createPostboxAddressRequestSigningAttributes{
			External: createPostboxAddressRequestSigningAttributesExternal{
				Selector:   req.Selector,
				PrivateKey: req.PrivateKey,
			},
		},
		LogOptions: createPostboxAddressRequestLogOptions{
			LogGroupID:          req.LogGroupID,
			IncludeMailStatuses: true,
		},
	}

	bodyBytes, err := y.sendHttpRequest(
		ctx,
		http.MethodPost,
		"https://console.yandex.cloud/gateway/root/postbox/createIdentity",
		reqStruct,
		req.YCUserCookie,
		req.YCUserCSRFToken,
	)

	if err != nil {
		return nil, err
	}

	resParsed := createPostboxAddressResponse{}
	err = json.Unmarshal(bodyBytes, &resParsed)

	if err != nil {
		return nil, errors.WithStack(err)
	}

	return &model.CreatePostboxAddressResponse{
		YCResponse: model.YCResponse{
			OperationID: resParsed.OperationID,
			Done:        resParsed.Done,
		},
		PostboxAddressID: resParsed.Response.ID,
	}, nil
}
