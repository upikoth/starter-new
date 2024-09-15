package config

import "github.com/kelseyhightower/envconfig"

type Config struct {
	GitHub GitHub
}

type GitHub struct {
	AccessToken string `envconfig:"GITHUB_ACCESS_TOKEN" required:"true"`
}

func New() (*Config, error) {
	config := &Config{}

	if err := envconfig.Process("", config); err != nil {
		return nil, err
	}

	return config, nil
}
