package yandexcloud

import (
	"context"
	"fmt"
	"net/http"
)

type updateServiceAccountRequest struct {
	AccessBindingDeltas []accessBindingDeltas `json:"accessBindingDeltas"`
}

type accessBindingDeltas struct {
	Action        string        `json:"action"`
	AccessBinding accessBinding `json:"accessBinding"`
}

type accessBinding struct {
	RoleID  string  `json:"roleId"`
	Subject subject `json:"subject"`
}

type subject struct {
	ID   string `json:"id"`
	Type string `json:"type"`
}

func (y *YandexCloud) UpdateServiceAccountRoles(
	ctx context.Context,
	accountID string,
	folderID string,
	roles []string,
) error {
	reqStruct := updateServiceAccountRequest{}

	for i := range roles {
		reqStruct.AccessBindingDeltas = append(reqStruct.AccessBindingDeltas, accessBindingDeltas{
			Action: "ADD",
			AccessBinding: accessBinding{
				RoleID: roles[i],
				Subject: subject{
					ID:   accountID,
					Type: "serviceAccount",
				},
			},
		})
	}

	_, err := y.sendHttpRequest(
		ctx,
		http.MethodPost,
		fmt.Sprintf("https://resource-manager.api.cloud.yandex.net/resource-manager/v1/folders/%s:updateAccessBindings", folderID),
		reqStruct,
	)

	if err != nil {
		return err
	}

	return nil
}
