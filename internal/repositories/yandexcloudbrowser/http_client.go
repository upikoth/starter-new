package yandexcloudbrowser

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"github.com/pkg/errors"
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
	y.rl.Take()

	body, err := json.Marshal(req)
	if err != nil {
		return []byte{}, errors.WithStack(err)
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
		return []byte{}, errors.WithStack(err)
	}

	client := &http.Client{}
	res, err := client.Do(request)
	if err != nil {
		return []byte{}, errors.WithStack(err)
	}

	defer func(body io.ReadCloser) {
		_ = body.Close()
	}(res.Body)

	bodyBytes, err := io.ReadAll(res.Body)
	if err != nil {
		return []byte{}, errors.WithStack(err)
	}

	if res.StatusCode != http.StatusOK {
		err := errors.New(fmt.Sprintf("не удалось выполнить запрос %s: %s, статус ответа - %d", method, url, res.StatusCode))
		y.logger.Error(err.Error())
		y.logger.Error(string(bodyBytes))
		return []byte{}, errors.WithStack(err)
	}

	return bodyBytes, nil
}
