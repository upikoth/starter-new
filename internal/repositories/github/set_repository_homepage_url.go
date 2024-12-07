package github

import (
	"context"
	"fmt"
	"github.com/upikoth/starter-new/internal/model"
	"net/http"
)

type setRepositoryHomepageURLRequest struct {
	Homepage string `json:"homepage"`
}

func (g *Github) SetRepositoryHomepageURL(ctx context.Context, req model.SetGithubRepositoryHomepageURLRequest) error {
	reqStruct := setRepositoryHomepageURLRequest{
		Homepage: req.URL,
	}

	_, err := g.sendHttpRequest(
		ctx,
		http.MethodPatch,
		fmt.Sprintf(
			"https://api.github.com/repos/%s/%s",
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
