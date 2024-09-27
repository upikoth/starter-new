package github

import (
	"context"
	"net/http"
)

type createRepositoryRequest struct {
	Name string `json:"name"`
}

func (r *Github) CreateRepository(ctx context.Context, repoName string) error {
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
		return err
	}

	return nil
}
