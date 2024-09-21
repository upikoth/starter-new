package github

import (
	"context"
	"fmt"
	"net/http"
)

type createRepositoryRequest struct {
	Name string `json:"name"`
}

func (r *Github) CreateRepository(ctx context.Context, repoName string) error {
	r.logger.Info(fmt.Sprintf("Создаем репозиторий в github: %s", repoName))

	reqStruct := createRepositoryRequest{
		Name: repoName,
	}

	_, err := r.sendHttpRequest(
		ctx,
		http.MethodPost,
		"https://api.github.com/user/repos",
		reqStruct,
	)

	if err != nil {
		r.logger.Error(err.Error())
		return err
	}

	r.logger.Info(fmt.Sprintf("Репозиторий в github успешно создан: %s", repoName))
	return nil
}
