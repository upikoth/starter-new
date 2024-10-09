package github

import (
	"context"
	"fmt"
	"github.com/upikoth/starter-new/internal/model"
	"net/http"
)

type addRepositoryVariableRequest struct {
	Name  string `json:"name"`
	Value string `json:"value"`
}

func (g *Github) AddRepositoryVariable(ctx context.Context, req model.AddGithubRepositoryVariableRequest) error {
	reqStruct := addRepositoryVariableRequest{
		Name:  req.VariableName,
		Value: req.VariableValue,
	}

	_, err := g.sendHttpRequest(
		ctx,
		http.MethodPost,
		fmt.Sprintf(
			"https://api.github.com/repos/%s/%s/actions/variables",
			req.GithubUserName,
			req.GithubRepoName,
		),
		reqStruct,
	)

	if err != nil {
		return err
	}

	return nil
}
