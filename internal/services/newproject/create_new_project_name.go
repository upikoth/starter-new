package newproject

import (
	"context"
	"fmt"
)

func (p *Service) CreateNewProjectName(ctx context.Context) error {
	fmt.Println("Введите название нового приложения. Латиница, строчные буквы, разделитель - дефис.")

	name, err := p.repositories.ConsoleInput.GetString(ctx)

	if err != nil {
		return err
	}

	p.newProject.SetName(name)
	p.logger.Info(fmt.Sprintf("Название нового приложения - %s", name))

	return nil
}
