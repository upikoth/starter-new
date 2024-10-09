package github

import (
	"context"
	"fmt"
	"github.com/upikoth/starter-new/internal/model"
	"net/http"
)

type addEnvironmentSecretRequest struct {
	EncryptedValue string `json:"encrypted_value"`
	KeyId          string `json:"key_id"`
}

func (g *Github) AddEnvironmentSecret(ctx context.Context, req model.AddGithubEnvironmentSecretRequest) error {
	reqStruct := addEnvironmentSecretRequest{
		EncryptedValue: req.VariableEncryptedValue,
		KeyId:          req.RepositoryPublicKeyID,
	}

	_, err := g.sendHttpRequest(
		ctx,
		http.MethodPut,
		fmt.Sprintf(
			"https://api.github.com/repos/%s/%s/environments/%s/secrets/%s",
			req.GithubUserName,
			req.GithubRepoName,
			req.EnvironmentName,
			req.VariableName,
		),
		reqStruct,
	)

	if err != nil {
		return err
	}

	return nil
}
