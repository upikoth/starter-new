package newproject

import (
	"context"
	"fmt"
	"github.com/pkg/errors"
	"golang.org/x/sync/errgroup"
	"os"
)

func (p *Service) CreateFrontendLocalDevEnvironment(ctx context.Context) error {
	eg, _ := errgroup.WithContext(ctx)

	eg.Go(func() error {
		file, err := os.Create(fmt.Sprintf(
			"%s/public/environment.json",
			p.newProject.GetFrontendLocalPath(),
		))

		if err != nil {
			return err
		}

		defer func() {
			_ = file.Close()
		}()

		_, err = file.WriteString(
			fmt.Sprintf(
				`{
  "API_URL": "%s",
  "SENTRY_DSN": "%s",
  "ENVIRONMENT": "development"
}
`,
				p.newProject.GetDomainURL(),
				p.newProject.GetSentryFrontendDSN(),
			))

		return err
	})

	eg.Go(func() error {
		file, err := os.Create(fmt.Sprintf(
			"%s/public/environment.json.local",
			p.newProject.GetFrontendLocalPath(),
		))

		if err != nil {
			return err
		}

		defer func() {
			_ = file.Close()
		}()

		_, err = file.WriteString(fmt.Sprintf(
			`API_URL=%s
SENTRY_DSN=%s
ENVIRONMENT=development

// Для деплоя приложения в object storage
S3_ACCESS_KEY_ID=%s
S3_SECRET_ACCESS_KEY=%s
S3_BUCKET_NAME=%s

// Для загрузки зависимостей из моего репозитория
// Нужно установить в переменные окружения при локальной разработке
UPIKOTH_PACKAGES_READ=%s

// Токен бота телеграм.
NOTIFICATIONS_TELEGRAM_TOKEN=%s
// Id бота телеграм.
NOTIFICATIONS_TELEGRAM_TO=%s
`,
			p.newProject.GetDomainURL(),
			p.newProject.GetSentryFrontendDSN(),
			p.newProject.GetYCObjectStorageAccessKeyID(),
			p.newProject.GetYCObjectStorageAccessKeySecret(),
			p.newProject.GetYCObjectStorageBucketNameStatic(),
			p.config.ProxyVariables.UpikothPackagesRead,
			p.config.ProxyVariables.NotificationsTelegramToken,
			p.config.ProxyVariables.NotificationsTelegramTo,
		))

		return nil
	})

	if err := eg.Wait(); err != nil {
		return errors.WithStack(err)
	}

	p.logger.Info("Local: локальные файлы для frontend с env переменными созданы")

	return nil
}
