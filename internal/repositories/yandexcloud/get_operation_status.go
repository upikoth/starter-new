package yandexcloud

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/pkg/errors"
	"net/http"
	"time"
)

type operationStatusResponse struct {
	Done bool `json:"done"`
}

func (y *YandexCloud) GetOperationStatus(ctx context.Context, operationID string) bool {
	for i := 0; i < 30; i += 1 {
		time.Sleep(time.Second)
		done, _ := y.getOperationStatus(ctx, operationID)

		if done {
			return true
		}
	}

	return false
}

func (y *YandexCloud) getOperationStatus(ctx context.Context, operationID string) (bool, error) {
	bodyBytes, err := y.sendHttpRequest(
		ctx,
		http.MethodGet,
		fmt.Sprintf("https://operation.api.cloud.yandex.net/operations/%s", operationID),
		struct{}{},
	)

	if err != nil {
		return false, err
	}

	resParsed := operationStatusResponse{}
	err = json.Unmarshal(bodyBytes, &resParsed)

	if err != nil {
		return false, errors.WithStack(err)
	}

	return resParsed.Done, nil
}
