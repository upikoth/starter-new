package yandexcloudbrowser

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"io"
	"net/http"
	"strconv"
	"time"
)

func (y *YandexCloudBrowser) sendHttpRequest(
	ctx context.Context,
	method string,
	url string,
	req any,
	cookie string,
	csrfToken string,
) ([]byte, error) {
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

	request.Header.Add("Cookie", cookie)
	request.Header.Add("X-CSRF-TOKEN", csrfToken)
	request.Header.Add("Host", "https://console.yandex.cloud")
	request.Header.Add("Content-Type", "application/json")
	request.Header.Add("Content-Length", strconv.Itoa(len(body)))

	if err != nil {
		return []byte{}, err
	}

	client := &http.Client{}
	res, err := client.Do(request)
	if err != nil {
		return []byte{}, err
	}

	defer res.Body.Close()

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
