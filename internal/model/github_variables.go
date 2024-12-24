package model

type BackendRepositoryVariables struct {
	SentryDSN              string `json:"SENTRY_DSN"`
	YcContainterName       string `json:"YC_CONTAINER_NAME"`
	YcFolderID             string `json:"YC_FOLDER_ID"`
	YcLogOptionsLogGroupID string `json:"YC_LOG_OPTIONS_LOG_GROUP_ID"`
	YcRegistry             string `json:"YC_REGISTRY"`
	YcServiceAccountID     string `json:"YC_SERVICE_ACCOUNT_ID"`
	OauthMailAPIURL        string `json:"OAUTH_MAIL_API_URL"`
	OauthYandexAPIURL      string `json:"OAUTH_YANDEX_API_URL"`
}

type BackendEnvironmentVariables struct {
	Environment                                 string `json:"ENVIRONMENT"`
	FrontConfirmationPasswordRecoveryRequestURL string `json:"FRONT_CONFIRMATION_PASSWORD_RECOVERY_REQUEST_URL"`
	FrontConfirmationRegistrationURL            string `json:"FRONT_CONFIRMATION_REGISTRATION_URL"`
	FrontURL                                    string `json:"FRONT_URL"`
	FrontHandleAuthPageURL                      string `json:"FRONT_HANDLE_AUTH_PAGE_URL"`
	Port                                        string `json:"PORT"`
	YCPFromAddress                              string `json:"YCP_FROM_ADDRESS"`
	YCPFromName                                 string `json:"YCP_FROM_NAME"`
	YCPHost                                     string `json:"YCP_HOST"`
	YCPPort                                     string `json:"YCP_PORT"`
	YCS3Path                                    string `json:"YC_S3_PATH"`
	YDBAuthFileDirName                          string `json:"YDB_AUTH_FILE_DIR_NAME"`
	YDBAuthFile                                 string `json:"YDB_AUTH_FILE_NAME"`
	OauthMailRedirectURL                        string `json:"OAUTH_MAIL_REDIRECT_URL"`
	OauthVKRedirectURL                          string `json:"OAUTH_VK_REDIRECT_URL"`
	OauthYandexRedirectURL                      string `json:"OAUTH_YANDEX_REDIRECT_URL"`
}

type FrontendRepositoryVariables struct {
	SentryDSN string `json:"SENTRY_DSN"`
}

type FrontendEnvironmentVariables struct {
	Environment  string `json:"ENVIRONMENT"`
	APIURL       string `json:"API_URL"`
	S3BucketName string `json:"S3_BUCKET_NAME"`
}
