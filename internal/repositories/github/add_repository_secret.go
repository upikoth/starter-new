package github

import (
	"context"
	"fmt"
	"github.com/upikoth/starter-new/internal/model"
	"net/http"
)

type addRepositorySecretRequest struct {
	EncryptedValue string `json:"encrypted_value"`
	KeyId          string `json:"key_id"`
}

func (g *Github) AddRepositorySecret(ctx context.Context, req model.AddGithubRepositorySecretRequest) error {
	reqStruct := addRepositorySecretRequest{
		EncryptedValue: req.VariableEncryptedValue,
		KeyId:          req.RepositoryPublicKeyID,
	}

	_, err := g.sendHttpRequest(
		ctx,
		http.MethodPut,
		fmt.Sprintf(
			"https://api.github.com/repos/%s/%s/actions/secrets/%s",
			req.GithubUserName,
			req.GithubRepoName,
			req.VariableName,
		),
		reqStruct,
	)

	if err != nil {
		return err
	}

	return nil
}
