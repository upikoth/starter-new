package yandexcloudbrowser

import (
	"context"
	"encoding/json"
	"net/http"

	"github.com/upikoth/starter-new/internal/model"
)

type createCertificateRequest struct {
	Req           string   `json:"*"`
	ChallengeType string   `json:"challengeType"`
	Domains       []string `json:"domains"`
	FolderID      string   `json:"folderId"`
	Name          string   `json:"name"`
	Provider      string   `json:"provider"`
}

type createCertificateResponse struct {
	OperationID string `json:"id"`
	Done        bool   `json:"done"`
	Metadata    struct {
		CertificateID string `json:"certificateId"`
	} `json:"metadata"`
}

func (y *YandexCloudBrowser) CreateCertificate(
	ctx context.Context,
	req model.CreateCertificateRequest,
) (*model.CreateCertificateResponse, error) {
	reqStruct := createCertificateRequest{
		Req:           "certificate-request",
		ChallengeType: "DNS",
		Domains:       []string{req.Domain},
		FolderID:      req.FolderID,
		Name:          req.CertificateName,
		Provider:      "LETS_ENCRYPT",
	}

	bodyBytes, err := y.sendHttpRequest(
		ctx,
		http.MethodPost,
		"https://console.yandex.cloud/gateway/root/certificateManager/requestNewCertificate",
		reqStruct,
		req.YCUserCookie,
		req.YCUserCSRFToken,
	)

	if err != nil {
		return nil, err
	}

	resParsed := createCertificateResponse{}
	err = json.Unmarshal(bodyBytes, &resParsed)

	if err != nil {
		return nil, err
	}

	return &model.CreateCertificateResponse{
		OperationID: resParsed.OperationID,
		Done:        resParsed.Done,
	}, nil
}
