package config

import (
	"github.com/kelseyhightower/envconfig"
	"github.com/pkg/errors"
)

type Config struct {
	GitHub            GitHub
	YandexCloud       YandexCloud
	ProxyVariables    ProxyVariables
	MainSiteDomain    string `envconfig:"MAIN_SITE_DOMAIN" required:"true"`
	ProjectsLocalPath string `envconfig:"PROJECTS_LOCAL_PATH" required:"true"`
}

type GitHub struct {
	AccessToken                string `envconfig:"GITHUB_ACCESS_TOKEN" required:"true"`
	UserName                   string `envconfig:"GITHUB_USER_NAME" required:"true"`
	BackendTemplateProjectName string `envconfig:"GITHUB_BACKEND_TEMPLATE_PROJECT_NAME" required:"true"`
}

type YandexCloud struct {
	OauthToken string `envconfig:"YANDEX_CLOUD_OAUTH_USER_TOKEN" required:"true"`
	CloudID    string `envconfig:"YANDEX_CLOUD_CLOUD_ID" required:"true"`
}

type ProxyVariables struct {
	NotificationsTelegramTo    string `envconfig:"NOTIFICATIONS_TELEGRAM_TO" required:"true"`
	NotificationsTelegramToken string `envconfig:"NOTIFICATIONS_TELEGRAM_TOKEN" required:"true"`
	UpikothPackagesRead        string `envconfig:"UPIKOTH_PACKAGES_READ" required:"true"`
}

func New() (*Config, error) {
	config := &Config{}

	if err := envconfig.Process("", config); err != nil {
		return nil, errors.WithStack(err)
	}

	return config, nil
}
