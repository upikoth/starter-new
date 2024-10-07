package newproject

import (
	"context"
	"fmt"
	"github.com/pkg/errors"
	"os"
	"os/exec"
)

func (p *Service) createLocalFrontendRepo(_ context.Context) error {
	dir, err := os.Getwd()

	if err != nil {
		return errors.WithStack(err)
	}

	_, err = exec.Command(
		"/bin/sh",
		fmt.Sprintf("%s/scripts/clone-starter-vue3-repo.sh", dir),
		p.newProject.GetLocalPath(),
		p.newProject.GetFrontendRepositoryName(),
		p.newProject.GetBackendRepositoryName(),
	).Output()

	if err != nil {
		return errors.WithStack(err)
	}

	return nil
}
