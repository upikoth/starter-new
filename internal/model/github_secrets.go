package model

type BackendRepositorySecrets struct {
	NotificationsTelegramTo    string `json:"NOTIFICATIONS_TELEGRAM_TO"`
	NotificationsTelegramToken string `json:"NOTIFICATIONS_TELEGRAM_TOKEN"`
	OauthMailClientID          string `json:"OAUTH_MAIL_CLIENT_ID"`
	OauthMailClientSecret      string `json:"OAUTH_MAIL_CLIENT_SECRET"`
	OauthVkClientID            string `json:"OAUTH_VK_CLIENT_ID"`
	OauthVkClientSecret        string `json:"OAUTH_VK_CLIENT_SECRET"`
	OauthYandexClientID        string `json:"OAUTH_YANDEX_CLIENT_ID"`
	OauthYandexClientSecret    string `json:"OAUTH_YANDEX_CLIENT_SECRET"`
}

type BackendEnvironmentSecrets struct {
	YCPostboxPassword   string `json:"YCP_PASSWORD"`
	YCPostboxUsername   string `json:"YCP_USERNAME"`
	YCSAJSONCredentials string `json:"YC_SA_JSON_CREDENTIALS"`
	YDBDSN              string `json:"YDB_DSN"`
}

type FrontendRepositorySecrets struct {
	NotificationsTelegramTo    string `json:"NOTIFICATIONS_TELEGRAM_TO"`
	NotificationsTelegramToken string `json:"NOTIFICATIONS_TELEGRAM_TOKEN"`
	UpikothPackagesRead        string `json:"UPIKOTH_PACKAGES_READ"`
}

type FrontendEnvironmentSecrets struct {
	YCObjectStorageAccessKeyID     string `json:"S3_ACCESS_KEY_ID"`
	YCObjectStorageSecretAccessKey string `json:"S3_SECRET_ACCESS_KEY"`
}
