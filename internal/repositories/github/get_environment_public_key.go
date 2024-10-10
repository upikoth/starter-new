package github

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/pkg/errors"
	"github.com/upikoth/starter-new/internal/model"
	"net/http"
)

type getEnvironmentPublicKeyResponse struct {
	KeyID string `json:"key_id"`
	Key   string `json:"key"`
}

func (g *Github) GetEnvironmentPublicKey(ctx context.Context, req model.GetGithubEnvironmentPublicKeyRequest) (*model.GithubEnvironmentPublicKey, error) {
	bodyBytes, err := g.sendHttpRequest(
		ctx,
		http.MethodGet,
		fmt.Sprintf(
			"https://api.github.com/repos/%s/%s/environments/%s/secrets/public-key",
			req.GithubUserName,
			req.GithubRepoName,
			req.EnvironmentName,
		),
		struct{}{},
	)

	if err != nil {
		return nil, err
	}

	resParsed := getEnvironmentPublicKeyResponse{}
	err = json.Unmarshal(bodyBytes, &resParsed)

	if err != nil {
		return nil, errors.WithStack(err)
	}

	return &model.GithubEnvironmentPublicKey{
		Key:   resParsed.Key,
		KeyID: resParsed.KeyID,
	}, nil
}
