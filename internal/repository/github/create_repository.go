package github

import (
	"bytes"
	"context"
	"errors"
	"fmt"
	"net/http"
	"time"
)

func (r *Github) CreateRepository(ctx context.Context, repoName string) error {
	r.logger.Info(fmt.Sprintf("Создаем репозиторий в github: %s", repoName))

	body := []byte(fmt.Sprintf(`{"name": "%s"}`, repoName))
	newCtx, cancel := context.WithTimeout(ctx, time.Minute)
	defer cancel()

	req, err := http.NewRequestWithContext(newCtx, http.MethodPost, "https://api.github.com/user/repos", bytes.NewBuffer(body))
	if err != nil {
		r.logger.Error(err.Error())
		return err
	}

	req.Header.Add("Authorization", fmt.Sprintf("Bearer %s", r.config.GitHub.AccessToken))
	req.Header.Add("X-GitHub-Api-Version", "2022-11-28")
	req.Header.Add("Accept", "application/vnd.github+json")

	client := &http.Client{}
	res, err := client.Do(req)
	if err != nil {
		r.logger.Error(err.Error())
		return err
	}

	defer res.Body.Close()

	if res.StatusCode != http.StatusCreated {
		err := errors.New("Репозиторий go не создан в github")
		r.logger.Error(err.Error())
		return err
	}

	r.logger.Info(fmt.Sprintf("Репозиторий в github успешно создан: %s", repoName))
	return nil
}
