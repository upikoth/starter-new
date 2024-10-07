package newproject

import (
	"context"
	"fmt"
	"os"
	"os/exec"
)

func (p *Service) createLocalBackendRepo(ctx context.Context) error {
	// 1. Создаем копию проекта starter-go.
	if err := p.cloneBackendTemplateProject(ctx); err != nil {
		return err
	}

	// 2. Перемещаем в правильную папку в проектах.
	if err := p.moveBackendToCorrectLocalFolder(ctx); err != nil {
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
		fmt.Sprintf("github.com/%s/%s", p.config.GitHub.UserName, p.newProject.GetBackendRepositoryName()),
	).Output()

	if err != nil {
		return err
	}

	return nil
}

func (p *Service) moveBackendToCorrectLocalFolder(_ context.Context) error {
	dir, err := os.Getwd()

	if err != nil {
		return err
	}

	_, err = exec.Command(
		"/bin/sh",
		fmt.Sprintf("%s/scripts/move-dir.sh", dir),
		fmt.Sprintf("%s/%s", dir, p.newProject.GetBackendRepositoryName()),
		p.newProject.GetLocalPath(),
	).Output()

	if err != nil {
		return err
	}

	return nil
}
