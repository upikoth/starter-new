package yandexcloud

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"time"
)

func (y *YandexCloud) sendHttpPostRequest(ctx context.Context, url string, req any) ([]byte, error) {
	y.logger.Info("Делаем POST запрос в YandexCloud")

	body, err := json.Marshal(req)
	if err != nil {
		y.logger.Error(err.Error())
		return []byte{}, err
	}

	ctxWithTimeout, cancel := context.WithTimeout(ctx, time.Minute)
	defer cancel()

	request, err := http.NewRequestWithContext(
		ctxWithTimeout,
		http.MethodPost,
		url,
		bytes.NewBuffer(body),
	)

	request.Header.Add("Authorization", fmt.Sprintf("Bearer %s", y.iamToken))

	if err != nil {
		y.logger.Error(err.Error())
		return []byte{}, err
	}

	client := &http.Client{}
	res, err := client.Do(request)
	if err != nil {
		y.logger.Error(err.Error())
		return []byte{}, err
	}

	defer res.Body.Close()

	bodyBytes, err := io.ReadAll(res.Body)
	if err != nil {
		y.logger.Error(err.Error())
		return []byte{}, err
	}

	if res.StatusCode != http.StatusOK {
		err := errors.New("не удалось выполнить POST запрос, статус ответа не 200")
		y.logger.Error(err.Error())
		y.logger.Error(string(bodyBytes))
		return []byte{}, err
	}

	return bodyBytes, nil
}
