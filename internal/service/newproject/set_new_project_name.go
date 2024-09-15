package newproject

import "fmt"

func (p *NewProject) SetNewProjectName() error {
	p.logger.Debug("Задаем имя нового проекта")

	name, err := p.repository.ConsoleInput.GetNewProjectName()

	if err != nil {
		p.logger.Error("Не удалось задать имя проекта")
		return err
	}

	p.project.Name = name

	p.logger.Info(fmt.Sprintf("Имя нового проекта: %s", name))
	return nil
}
