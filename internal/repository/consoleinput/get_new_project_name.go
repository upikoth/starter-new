package consoleinput

import (
	"context"
	"fmt"
)

func (c *ConsoleInput) GetNewProjectName(_ context.Context) (string, error) {
	c.logger.Debug("Запрашиваем имя нового проекта")

	projectName := ""

	fmt.Println("Введите название нового приложения. Латиница, строчные буквы, разделитель - дефис.")
	if _, err := fmt.Scanln(&projectName); err != nil {
		c.logger.Error(err.Error())
		return "", err
	}

	return projectName, nil
}
