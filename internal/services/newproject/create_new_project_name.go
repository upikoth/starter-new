package newproject

import (
	"context"
	"fmt"
	"regexp"
)

func (p *Service) CreateNewProjectName(ctx context.Context) error {
	fmt.Println("Введите название нового приложения. Латиница, строчные буквы, разделитель - дефис.")

	name := ""
	isNameValid := false

	for !isNameValid {
		var err error
		name, err = p.repositories.ConsoleInput.GetString(ctx)

		if err != nil {
			return err
		}

		if checkIsProjectNameValid(name) {
			isNameValid = true
		} else {
			fmt.Println("Введенное имя не соответствует формату. Латиница, строчные буквы, разделитель - дефис.")
		}
	}

	p.newProject.SetName(name)
	p.logger.Info(fmt.Sprintf("Название нового приложения - %s", name))

	return nil
}

func checkIsProjectNameValid(projectName string) bool {
	isMatch, err := regexp.MatchString("^[a-z][a-z0-9-]*$", projectName)

	if err != nil {
		return false
	}

	return isMatch
}
