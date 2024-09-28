package github

import (
	"context"
	"net/http"
)

type createRepositoryRequest struct {
	Name string `json:"name"`
}

func (g *Github) CreateRepository(ctx context.Context, repoName string) error {
	reqStruct := createRepositoryRequest{
		Name: repoName,
	}

	_, err := g.sendHttpRequest(
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
