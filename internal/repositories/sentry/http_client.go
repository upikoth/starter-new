package sentry

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

func (s *Sentry) sendHttpRequest(ctx context.Context, method string, url string, req any) ([]byte, error) {
	s.rl.Take()

	body, err := json.Marshal(req)
	if err != nil {
		return []byte{}, errors.WithStack(err)
	}

	if ctx.Err() != nil {
		return []byte{}, errors.WithStack(ctx.Err())
	}
	ctxWithTimeout, cancel := context.WithTimeout(ctx, time.Minute)
	defer cancel()

	request, err := http.NewRequestWithContext(
		ctxWithTimeout,
		method,
		url,
		bytes.NewBuffer(body),
	)

	request.Header.Add("Authorization", fmt.Sprintf("Bearer %s", s.config.Sentry.AuthToken))
	request.Header.Add("Content-Type", "application/json")

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
		s.logger.Error(err.Error())
		s.logger.Error(string(bodyBytes))
		return []byte{}, err
	}

	return bodyBytes, nil
}
