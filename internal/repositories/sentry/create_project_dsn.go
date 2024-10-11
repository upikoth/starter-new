package sentry

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/pkg/errors"
	"github.com/upikoth/starter-new/internal/model"
	"net/http"
)

type createProjectDSNResponse struct {
	Id  string `json:"id"`
	Dsn struct {
		Public string `json:"public"`
	} `json:"dsn"`
}

func (s *Sentry) CreateProjectDSN(ctx context.Context, req model.CreateSentryProjectDSNRequest) (string, error) {
	bodyBytes, err := s.sendHttpRequest(
		ctx,
		http.MethodPost,
		fmt.Sprintf("https://sentry.io/api/0/projects/%s/%s/keys/", s.config.Sentry.OrganizationID, req.ProjectName),
		struct{}{},
	)

	if err != nil {
		return "", err
	}

	resParsed := createProjectDSNResponse{}
	err = json.Unmarshal(bodyBytes, &resParsed)

	if err != nil {
		return "", errors.WithStack(err)
	}

	return resParsed.Dsn.Public, nil
}
