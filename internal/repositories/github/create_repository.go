package github

import (
	"context"
	"encoding/json"
	"github.com/pkg/errors"
	"net/http"
)

type createRepositoryRequest struct {
	Name string `json:"name"`
}

type createRepositoryResponse struct {
	ID int `json:"id"`
}

func (g *Github) CreateRepository(ctx context.Context, repoName string) (repositoryID int, err error) {
	reqStruct := createRepositoryRequest{
		Name: repoName,
	}

	bodyBytes, err := g.sendHttpRequest(
		ctx,
		http.MethodPost,
		"https://api.github.com/user/repos",
		reqStruct,
	)

	if err != nil {
		return 0, err
	}

	resParsed := createRepositoryResponse{}
	err = json.Unmarshal(bodyBytes, &resParsed)

	if err != nil {
		return 0, errors.WithStack(err)
	}

	return resParsed.ID, nil
}
