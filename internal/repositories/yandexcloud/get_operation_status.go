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

func (y *YandexCloud) GetOperationStatus(ctx context.Context, operationID string) (bool, error) {
	for i := 0; i < 120; i += 1 {
		time.Sleep(time.Second * 5)
		done, err := y.getOperationStatus(ctx, operationID)

		if err != nil {
			return false, err
		}

		if done {
			return true, nil
		}
	}

	return false, nil
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
