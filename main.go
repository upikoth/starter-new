package main

import (
	"github.com/joho/godotenv"
	"github.com/upikoth/starter-new/internal/app"
	"github.com/upikoth/starter-new/internal/config"
	"github.com/upikoth/starter-new/internal/pkg/logger"
)

func main() {
	_ = godotenv.Load()
	loggerInstance := logger.New()
	loggerInstance.SetPrettyOutputToConsole()

	config, err := config.New()
	if err != nil {
		loggerInstance.Fatal(err.Error())
	}

	app, err := app.New(config, loggerInstance)
	if err != nil {
		loggerInstance.Fatal(err.Error())
	}

	loggerInstance.Info("Запуск приложения")

	if appErr := app.Start(); appErr != nil {
		loggerInstance.Error("Приложение отработало с ошибкой")
	} else {
		loggerInstance.Info("Приложение успешно завершило работу")
	}
}
