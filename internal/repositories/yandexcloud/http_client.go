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

func (y *YandexCloud) sendHttpRequest(ctx context.Context, method string, url string, req any) ([]byte, error) {
	body, err := json.Marshal(req)
	if err != nil {
		return []byte{}, err
	}

	ctxWithTimeout, cancel := context.WithTimeout(ctx, time.Minute)
	defer cancel()

	request, err := http.NewRequestWithContext(
		ctxWithTimeout,
		method,
		url,
		bytes.NewBuffer(body),
	)

	request.Header.Add("Authorization", fmt.Sprintf("Bearer %s", y.iamToken))

	if err != nil {
		return []byte{}, err
	}

	client := &http.Client{}
	res, err := client.Do(request)
	if err != nil {
		return []byte{}, err
	}

	defer func(body io.ReadCloser) {
		_ = body.Close()
	}(res.Body)

	bodyBytes, err := io.ReadAll(res.Body)
	if err != nil {
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
