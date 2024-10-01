package yandexcloudbrowser

import (
	"context"
	"encoding/json"
	"github.com/upikoth/starter-new/internal/model"
	"net/http"
)

type getPostboxVerificationRecordsResponse struct {
	Records []struct {
		Type  string `json:"type"`
		Name  string `json:"name"`
		Value string `json:"value"`
	} `json:"records"`
}

type getPostboxVerificationRequest struct {
	IdentityID string `json:"identityId"`
}

func (y *YandexCloudBrowser) GetPostboxVerificationRecord(
	ctx context.Context,
	req model.GetPostboxVerificationRecordRequest,
) (*model.PostboxVerificationRecord, error) {
	reqStruct := getPostboxVerificationRequest{
		IdentityID: req.IdentityID,
	}

	bodyBytes, err := y.sendHttpRequest(
		ctx,
		http.MethodPost,
		"https://console.yandex.cloud/gateway/root/postbox/getIdentityVerificationRecords",
		reqStruct,
		req.YCUserCookie,
		req.YCUserCSRFToken,
	)

	if err != nil {
		return nil, err
	}

	resParsed := getPostboxVerificationRecordsResponse{}
	err = json.Unmarshal(bodyBytes, &resParsed)

	if err != nil {
		return nil, err
	}

	record := model.PostboxVerificationRecord{
		Type:  resParsed.Records[0].Type,
		Name:  resParsed.Records[0].Name,
		Value: resParsed.Records[0].Value,
	}

	return &record, nil
}
