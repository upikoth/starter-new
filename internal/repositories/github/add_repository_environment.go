package github

import (
	"context"
	"fmt"
	"github.com/upikoth/starter-new/internal/model"
	"net/http"
)

func (g *Github) AddRepositoryEnvironment(ctx context.Context, req model.AddGithubRepositoryEnvironmentRequest) error {
	_, err := g.sendHttpRequest(
		ctx,
		http.MethodPut,
		fmt.Sprintf(
			"https://api.github.com/repos/%s/%s/environments/%s",
			req.GithubUserName,
			req.GithubRepoName,
			req.EnvironmentName,
		),
		struct{}{},
	)

	if err != nil {
		return err
	}

	return nil
}
