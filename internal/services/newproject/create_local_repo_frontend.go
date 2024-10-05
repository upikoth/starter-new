package newproject

import (
	"context"
	"fmt"
	"os"
	"os/exec"
)

func (p *Service) createFrontendRepo(ctx context.Context) error {
	dir, err := os.Getwd()

	if err != nil {
		return err
	}

	_, err = exec.Command(
		"/bin/sh",
		fmt.Sprintf("%s/scripts/clone-starter-vue3-repo.sh", dir),
		p.getProjectLocalPath(),
		p.getFrontendRepoName(),
		p.getBackendRepoName(),
	).Output()

	if err != nil {
		return err
	}

	return nil
}
