package newproject

import (
	"context"
	"fmt"
	"github.com/pkg/errors"
	"golang.org/x/sync/errgroup"
	"os"
	"strings"
)

func (p *Service) CreateBackendLocalDevEnvironment(ctx context.Context) error {
	eg, _ := errgroup.WithContext(ctx)

	eg.Go(func() error {
		err := os.Mkdir(
			fmt.Sprintf("%s%s", p.newProject.GetBackendLocalPath(), p.newProject.GetYCYDBFileDirName()),
			0777,
		)

		if err != nil {
			return err
		}

		file, err := os.Create(fmt.Sprintf(
			"%s%s/%s",
			p.newProject.GetBackendLocalPath(),
			p.newProject.GetYCYDBFileDirName(),
			p.newProject.GetYCYDBFileName()),
		)

		if err != nil {
			return err
		}

		defer func() {
			_ = file.Close()
		}()

		_, err = file.WriteString(p.newProject.GetYCSAJSONCredentials())

		return err
	})

	eg.Go(func() error {
		file, err := os.Create(fmt.Sprintf(
			"%s/%s",
			p.newProject.GetBackendLocalPath(),
			".env",
		))

		if err != nil {
			return err
		}

		defer func() {
			_ = file.Close()
		}()

		_, err = file.WriteString(fmt.Sprintf(
			`PORT=%s
ENVIRONMENT=development

YCP_HOST=%s
YCP_PORT=%s
YCP_FROM_NAME=%s
YCP_FROM_ADDRESS=%s
YCP_USERNAME=%s
YCP_PASSWORD=%s

YDB_DSN=%s
YDB_AUTH_FILE_DIR_NAME=%s
YDB_AUTH_FILE_NAME=%s
# Данные для авторизации в yandex cloud. Используются в ci/cd и при поключении к ydb.
YC_SA_JSON_CREDENTIALS=%s

FRONT_URL=%s
FRONT_CONFIRMATION_REGISTRATION_URL="%s"
FRONT_CONFIRMATION_PASSWORD_RECOVERY_REQUEST_URL="%s"

SENTRY_DSN=%s

# Токен бота телеграм.
# NOTIFICATIONS_TELEGRAM_TOKEN=%s
# Id бота телеграм.
# NOTIFICATIONS_TELEGRAM_TO=%s

# YC_CONTAINER_NAME=%s
# YC_REGISTRY=%s
# YC_FOLDER_ID=%s
# YC_SERVICE_ACCOUNT_ID=%s
# YC_LOG_OPTIONS_LOG_GROUP_ID=%s
# YC_S3_PATH=%s
`,
			p.newProject.GetBackendPort(),
			p.newProject.GetYCPHost(),
			p.newProject.GetYCPPort(),
			p.newProject.GetEmailFromName(),
			p.newProject.GetEmailFromAddress(),
			p.newProject.GetYCPostboxUsername(),
			p.newProject.GetYCPostboxPassword(),
			p.newProject.GetYCYDBEndpoint(),
			strings.Replace(p.newProject.GetYCYDBFileDirName(), "/", "", 1),
			p.newProject.GetYCYDBFileName(),
			p.newProject.GetYCSAJSONCredentials(),
			p.newProject.GetDomain(),
			p.newProject.GetFrontendConfirmationRegistrationURL(),
			p.newProject.GetFrontendConfirmationPasswordRecoveryRequestURL(),
			p.newProject.GetSentryBackendDSN(),
			p.config.ProxyVariables.NotificationsTelegramToken,
			p.config.ProxyVariables.NotificationsTelegramTo,
			p.newProject.GetYCServerlessContainerName(),
			p.newProject.GetYCContainerRegistryID(),
			p.newProject.GetYCFolderID(),
			p.newProject.GetYCServiceAccountID(),
			p.newProject.GetYCLoggingGroupID(),
			p.newProject.GetYCObjectStorageBucketNameSecrets(),
		))

		return nil
	})

	if err := eg.Wait(); err != nil {
		return errors.WithStack(err)
	}

	p.logger.Info("Local: локальные файлы для backend с env переменными созданы")

	return nil
}
