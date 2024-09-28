package main

import (
	"context"

	"github.com/joho/godotenv"
	"github.com/upikoth/starter-new/internal/app"
	"github.com/upikoth/starter-new/internal/config"
	"github.com/upikoth/starter-new/internal/pkg/logger"
)

func main() {
	_ = godotenv.Load()
	loggerInstance := logger.New()
	loggerInstance.SetPrettyOutputToConsole()

	cfg, err := config.New()
	if err != nil {
		loggerInstance.Fatal(err.Error())
	}

	appInstance, err := app.New(cfg, loggerInstance)
	if err != nil {
		loggerInstance.Fatal(err.Error())
	}

	loggerInstance.Info("Запуск приложения")
	ctx := context.Background()

	if appErr := appInstance.Start(ctx); appErr != nil {
		loggerInstance.Error("Приложение отработало с ошибкой")
	} else {
		loggerInstance.Info("Приложение успешно завершило работу")
	}
}
