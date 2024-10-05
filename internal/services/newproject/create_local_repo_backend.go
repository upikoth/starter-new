package newproject

import (
	"context"
	"fmt"
	"os"
	"os/exec"
)

func (p *Service) createBackendRepo(ctx context.Context) error {
	// 1. Создаем копию проекта starter-go.
	if err := p.cloneBackendTemplateProject(ctx); err != nil {
		return err
	}

	// 2. Перемещаем в правильную папку в проектах.
	if err := p.moveBackendToCorrectFolder(ctx); err != nil {
		return err
	}

	return nil
}

func (p *Service) cloneBackendTemplateProject(_ context.Context) error {
	dir, err := os.Getwd()

	if err != nil {
		return err
	}

	_, err = exec.Command(
		"/bin/sh",
		fmt.Sprintf("%s/scripts/clone-starter-go-repo.sh", dir),
		fmt.Sprintf("github.com/%s/%s", p.config.GitHub.UserName, p.config.GitHub.BackendTemplateProjectName),
		fmt.Sprintf("github.com/%s/%s", p.config.GitHub.UserName, p.getBackendRepoName()),
	).Output()

	if err != nil {
		return err
	}

	return nil
}

func (p *Service) moveBackendToCorrectFolder(_ context.Context) error {
	dir, err := os.Getwd()

	if err != nil {
		return err
	}

	_, err = exec.Command(
		"/bin/sh",
		fmt.Sprintf("%s/scripts/move-dir.sh", dir),
		fmt.Sprintf("%s/%s", dir, p.getBackendRepoName()),
		p.getProjectLocalPath(),
	).Output()

	if err != nil {
		return err
	}

	return nil
}
