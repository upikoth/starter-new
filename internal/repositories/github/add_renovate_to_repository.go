package github

import (
	"context"
	"fmt"
	"net/http"
)

const renovateInstallationID = "49360187"

func (g *Github) AddRenovateToRepository(ctx context.Context, repoID int) error {
	_, err := g.sendHttpRequest(
		ctx,
		http.MethodPut,
		fmt.Sprintf("https://api.github.com/user/installations/%s/repositories/%d", renovateInstallationID, repoID),
		struct{}{},
	)

	if err != nil {
		return err
	}

	return nil
}
