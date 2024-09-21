package yandexcloud

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"time"
)

type operationStatusResponse struct {
	Done bool `json:"done"`
}

func (y *YandexCloud) GetOperationStatus(ctx context.Context, operationID string) bool {
	for i := 0; i < 10; i += 1 {
		time.Sleep(time.Second)
		done, _ := y.getOperationStatus(ctx, operationID)

		if done {
			return true
		}
	}

	return false
}

func (y *YandexCloud) getOperationStatus(ctx context.Context, operationID string) (bool, error) {
	y.logger.Info("Проверяем статус операции")

	bodyBytes, err := y.sendHttpRequest(
		ctx,
		http.MethodGet,
		fmt.Sprintf("https://operation.api.cloud.yandex.net/operations/%s", operationID),
		struct{}{},
	)

	if err != nil {
		y.logger.Error(err.Error())
		return false, err
	}

	resParsed := operationStatusResponse{}
	err = json.Unmarshal(bodyBytes, &resParsed)

	if err != nil {
		y.logger.Error(err.Error())
		return false, err
	}

	y.logger.Info("Статус получен!")

	return resParsed.Done, nil
}
