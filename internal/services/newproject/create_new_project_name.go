package newproject

import (
	"context"
	"fmt"
)

func (p *NewProjectService) CreateNewProjectName(ctx context.Context) error {
	fmt.Println("Введите название нового приложения. Латиница, строчные буквы, разделитель - дефис.")

	name, err := p.repositories.ConsoleInput.GetString(ctx)

	if err != nil {
		return err
	}

	p.newProject.name = name

	return nil
}
