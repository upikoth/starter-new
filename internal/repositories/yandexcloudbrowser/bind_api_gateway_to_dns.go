package yandexcloudbrowser

import (
	"context"
	"github.com/upikoth/starter-new/internal/model"
	"net/http"
)

type bindApiGatewayToDNSRequest struct {
	Additions []bindApiGatewayToDNSRequestAddition `json:"additions"`
}

type bindApiGatewayToDNSRequestAddition struct {
	DnsZoneID   string                                  `json:"dnsZoneId"`
	Name        string                                  `json:"name"`
	Type        string                                  `json:"type"`
	Data        []string                                `json:"data"`
	Ttl         int                                     `json:"ttl"`
	UiProtected bool                                    `json:"uiProtected"`
	Owner       bindApiGatewayToDNSRequestAdditionOwner `json:"owner"`
}

type bindApiGatewayToDNSRequestAdditionOwner struct {
	OwnerType string `json:"ownerType"`
	OwnerId   string `json:"ownerId"`
}

func (y *YandexCloudBrowser) BindApiGatewayToDNS(
	ctx context.Context,
	req model.YCBindApiGatewayToDNSRequest,
) error {
	reqStruct := bindApiGatewayToDNSRequest{
		Additions: []bindApiGatewayToDNSRequestAddition{
			{
				Data: []string{
					req.DNSRecordText,
				},
				DnsZoneID: req.DNSZoneID,
				Name:      req.DNSRecordName,
				Owner: bindApiGatewayToDNSRequestAdditionOwner{
					OwnerId:   req.DNSRecordOwnerID,
					OwnerType: "SERVERLESS_API_GATEWAY",
				},
				Ttl:         600, // 10 минут.
				UiProtected: false,
				Type:        "ANAME",
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
