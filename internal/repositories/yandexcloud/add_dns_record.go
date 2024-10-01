package yandexcloud

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/upikoth/starter-new/internal/model"
)

type addDNSRecordRequest struct {
	Additions []addDNSRecordRequestAddition `json:"additions"`
}

type addDNSRecordRequestAddition struct {
	Name string   `json:"name"`
	Type string   `json:"type"`
	Ttl  string   `json:"ttl"`
	Data []string `json:"data"`
}

type addDNSRecordResponse struct {
	OperationID string `json:"id"`
	Done        bool   `json:"done"`
}

func (y *YandexCloud) AddDNSRecord(
	ctx context.Context,
	dnsZoneID string,
	record model.DNSRecord,
) (*model.AddDNSRecordResponse, error) {
	reqStruct := addDNSRecordRequest{
		Additions: []addDNSRecordRequestAddition{
			{
				Name: record.Name,
				Type: record.Type,
				Ttl:  "600", // 10 минут
				Data: []string{record.Value},
			},
		},
	}

	bodyBytes, err := y.sendHttpRequest(
		ctx,
		http.MethodPost,
		fmt.Sprintf("https://dns.api.cloud.yandex.net/dns/v1/zones/%s:updateRecordSets", dnsZoneID),
		reqStruct,
	)

	if err != nil {
		return nil, err
	}

	resParsed := addDNSRecordResponse{}
	err = json.Unmarshal(bodyBytes, &resParsed)

	if err != nil {
		return nil, err
	}

	return &model.AddDNSRecordResponse{
		OperationID: resParsed.OperationID,
		Done:        resParsed.Done,
	}, nil
}
