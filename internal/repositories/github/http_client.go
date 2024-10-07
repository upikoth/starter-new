package github

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"github.com/pkg/errors"
	"io"
	"net/http"
	"time"
)

func (g *Github) sendHttpRequest(ctx context.Context, method string, url string, req any) ([]byte, error) {
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

	request.Header.Add("Authorization", fmt.Sprintf("Bearer %s", g.config.GitHub.AccessToken))
	request.Header.Add("X-GitHub-Api-Version", "2022-11-28")
	request.Header.Add("Accept", "application/vnd.github+json")

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

	if res.StatusCode != http.StatusOK && res.StatusCode != http.StatusCreated {
		err := errors.New(fmt.Sprintf("не удалось выполнить запрос %s: %s, статус ответа - %d", method, url, res.StatusCode))
		g.logger.Error(err.Error())
		g.logger.Error(string(bodyBytes))
		return []byte{}, errors.WithStack(err)
	}

	return bodyBytes, nil
}
