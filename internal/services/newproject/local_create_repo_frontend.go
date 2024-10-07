package newproject

import (
	"context"
	"fmt"
	"os"
	"os/exec"
)

func (p *Service) createLocalFrontendRepo(_ context.Context) error {
	dir, err := os.Getwd()

	if err != nil {
		return err
	}

	_, err = exec.Command(
		"/bin/sh",
		fmt.Sprintf("%s/scripts/clone-starter-vue3-repo.sh", dir),
		p.newProject.GetLocalPath(),
		p.newProject.GetFrontendRepositoryName(),
		p.newProject.GetBackendRepositoryName(),
	).Output()

	if err != nil {
		return err
	}

	return nil
}
