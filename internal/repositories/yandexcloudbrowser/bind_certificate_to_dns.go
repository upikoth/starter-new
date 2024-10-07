package yandexcloudbrowser

import (
	"context"
	"github.com/upikoth/starter-new/internal/model"
	"net/http"
)

type bindCertificateToDNSRequest struct {
	Additions []bindCertificateToDNSRequestAddition `json:"additions"`
}

type bindCertificateToDNSRequestAddition struct {
	DnsZoneID   string                                   `json:"dnsZoneId"`
	Name        string                                   `json:"name"`
	Type        string                                   `json:"type"`
	Data        []string                                 `json:"data"`
	Ttl         int                                      `json:"ttl"`
	UiProtected bool                                     `json:"uiProtected"`
	Owner       bindCertificateToDNSRequestAdditionOwner `json:"owner"`
}

type bindCertificateToDNSRequestAdditionOwner struct {
	OwnerType string `json:"ownerType"`
	OwnerId   string `json:"ownerId"`
}

func (y *YandexCloudBrowser) BindCertificateToDNS(
	ctx context.Context,
	req model.YCBindCertificateToDNSRequest,
) error {
	reqStruct := bindCertificateToDNSRequest{
		Additions: []bindCertificateToDNSRequestAddition{
			{
				Data: []string{
					req.DNSRecordText,
				},
				DnsZoneID: req.DNSZoneID,
				Name:      req.DNSRecordName,
				Owner: bindCertificateToDNSRequestAdditionOwner{
					OwnerId:   req.DNSRecordOwnerID,
					OwnerType: "CERTIFICATE_MANAGER_CERTIFICATE",
				},
				Ttl:         600, // 10 минут.
				UiProtected: false,
				Type:        "CNAME",
			},
		},
	}

	_, err := y.sendHttpRequest(
		ctx,
		http.MethodPost,
		"https://console.yandex.cloud/gateway/root/dns/addOwnedRecordSets",
		reqStruct,
		req.YCUserCookie,
		req.YCUserCSRFToken,
	)

	return err
}
