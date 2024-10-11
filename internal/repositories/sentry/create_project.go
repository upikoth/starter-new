package sentry

import (
	"context"
	"fmt"
	"github.com/upikoth/starter-new/internal/model"
	"net/http"
)

type createProjectRequest struct {
	Name     string `json:"name"`
	Platform string `json:"platform"`
}

func (s *Sentry) CreateProject(ctx context.Context, req model.CreateSentryProjectRequest) error {
	reqStruct := createProjectRequest{
		Name:     req.ProjectName,
		Platform: req.ProjectPlatform,
	}

	_, err := s.sendHttpRequest(
		ctx,
		http.MethodPost,
		fmt.Sprintf("https://sentry.io/api/0/teams/%s/%s/projects/", s.config.Sentry.OrganizationID, s.config.Sentry.TeamID),
		reqStruct,
	)

	if err != nil {
		return err
	}

	return nil
}
