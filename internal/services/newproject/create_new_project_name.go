package newproject

import (
	"context"
	"fmt"
)

func (p *NewProject) CreateNewProjectName(ctx context.Context) error {
	p.logger.Debug("Задаем имя нового проекта")

	name, err := p.repositories.ConsoleInput.GetNewProjectName(ctx)

	if err != nil {
		p.logger.Error("Не удалось задать имя проекта")
		return err
	}

	p.project.Name = name

	p.logger.Info(fmt.Sprintf("Имя нового проекта: %s", name))
	return nil
}
