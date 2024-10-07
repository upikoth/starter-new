package yandexcloudbrowser

import (
	"context"
	"encoding/json"
	"github.com/pkg/errors"
	"github.com/upikoth/starter-new/internal/model"
	"net/http"
)

type getCertificateResponse struct {
	Challenges []struct {
		DnsName      string `json:"dnsName"`
		DnsTxt       string `json:"dnsTxt"`
		DnsChallenge struct {
			Type string `json:"type"`
		} `json:"dnsChallenge"`
	} `json:"challenges"`
}

type getCertificateRequest struct {
	CertificateID string `json:"certificateId"`
	View          string `json:"view"`
}

func (y *YandexCloudBrowser) GetCertificateChallenge(
	ctx context.Context,
	req model.YCGetCertificateChallengeRequest,
) (*model.CertificateChallenge, error) {
	reqStruct := getCertificateRequest{
		CertificateID: req.CertificateID,
		View:          "FULL",
	}

	bodyBytes, err := y.sendHttpRequest(
		ctx,
		http.MethodPost,
		"https://console.yandex.cloud/gateway/root/certificateManager/getCertificate",
		reqStruct,
		req.YCUserCookie,
		req.YCUserCSRFToken,
	)

	if err != nil {
		return nil, err
	}

	resParsed := getCertificateResponse{}
	err = json.Unmarshal(bodyBytes, &resParsed)

	if err != nil {
		return nil, errors.WithStack(err)
	}

	challenge := model.CertificateChallenge{}

	for _, challange := range resParsed.Challenges {
		if challange.DnsChallenge.Type == "CNAME" {
			challenge.DNSName = challange.DnsName
			challenge.DNSText = challange.DnsTxt
			challenge.ChallegeType = challange.DnsChallenge.Type
			break
		}
	}

	return &challenge, nil
}
