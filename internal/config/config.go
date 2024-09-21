package config

import "github.com/kelseyhightower/envconfig"

type Config struct {
	GitHub             GitHub
	YandexCloud        YandexCloud
	MainSiteDomainName string `envconfig:"MAIN_SITE_DOMAIN_NAME" required:"true"`
}

type GitHub struct {
	AccessToken string `envconfig:"GITHUB_ACCESS_TOKEN" required:"true"`
	UserName    string `envconfig:"GITHUB_USER_NAME" required:"true"`
}

type YandexCloud struct {
	OauthToken string `envconfig:"YANDEX_CLOUD_OAUTH_USER_TOKEN" required:"true"`
	CloudID    string `envconfig:"YANDEX_CLOUD_CLOUD_ID" required:"true"`
}

func New() (*Config, error) {
	config := &Config{}

	if err := envconfig.Process("", config); err != nil {
		return nil, err
	}

	return config, nil
}
