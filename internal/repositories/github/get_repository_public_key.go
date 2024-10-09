package github

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/pkg/errors"
	"github.com/upikoth/starter-new/internal/model"
	"net/http"
)

type getRepositoryPublicKeyResponse struct {
	KeyID string `json:"key_id"`
	Key   string `json:"key"`
}

func (g *Github) GetRepositoryPublicKey(ctx context.Context, req model.GetGithubRepositoryPublicKeyRequest) (*model.GithubRepositoryPublicKey, error) {
	bodyBytes, err := g.sendHttpRequest(
		ctx,
		http.MethodGet,
		fmt.Sprintf("https://api.github.com/repos/%s/%s/actions/secrets/public-key", req.GithubUserName, req.GithubRepoName),
		struct{}{},
	)

	if err != nil {
		return nil, err
	}

	resParsed := getRepositoryPublicKeyResponse{}
	err = json.Unmarshal(bodyBytes, &resParsed)

	if err != nil {
		return nil, errors.WithStack(err)
	}

	return &model.GithubRepositoryPublicKey{
		Key:   resParsed.Key,
		KeyID: resParsed.KeyID,
	}, nil
}
