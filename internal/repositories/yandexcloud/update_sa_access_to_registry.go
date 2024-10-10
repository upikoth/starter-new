package yandexcloud

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/pkg/errors"
	"github.com/upikoth/starter-new/internal/model"
	"net/http"
)

type updateServiceAccountAccessToRegistryRequest struct {
	AccessBindingDeltas []accessBindingToRegistryDeltas `json:"accessBindingDeltas"`
}

type accessBindingToRegistryDeltas struct {
	Action        string                  `json:"action"`
	AccessBinding accessBindingToRegistry `json:"accessBinding"`
}

type accessBindingToRegistry struct {
	RoleID  string               `json:"roleId"`
	Subject subjectAccessBinding `json:"subject"`
}

type subjectAccessBinding struct {
	ID   string `json:"id"`
	Type string `json:"type"`
}

type updateServiceAccountAccessTyRegistryResponse struct {
	OperationID string `json:"id"`
	Done        bool   `json:"done"`
}

func (y *YandexCloud) UpdateServiceAccountAccessToRegistry(
	ctx context.Context,
	accountID string,
	registryID string,
) (*model.UpdateServiceAccountAccessToRegistryResponse, error) {
	reqStruct := updateServiceAccountAccessToRegistryRequest{
		[]accessBindingToRegistryDeltas{
			{
				Action: "ADD",
				AccessBinding: accessBindingToRegistry{
					RoleID: "editor",
					Subject: subjectAccessBinding{
						ID:   accountID,
						Type: "serviceAccount",
					},
				},
			},
		},
	}

	bodyBytes, err := y.sendHttpRequest(
		ctx,
		http.MethodPost,
		fmt.Sprintf("https://container-registry.api.cloud.yandex.net/container-registry/v1/registries/%s:updateAccessBindings", registryID),
		reqStruct,
	)

	if err != nil {
		return nil, err
	}

	resParsed := updateServiceAccountAccessTyRegistryResponse{}
	err = json.Unmarshal(bodyBytes, &resParsed)

	if err != nil {
		return nil, errors.WithStack(err)
	}

	return &model.UpdateServiceAccountAccessToRegistryResponse{
		YCResponse: model.YCResponse{
			OperationID: resParsed.OperationID,
			Done:        resParsed.Done,
		},
	}, nil
}
