package model

type YCBrowserRequest struct {
	YCUserCookie    string
	YCUserCSRFToken string
}

type YCCreateCertificateRequest struct {
	YCBrowserRequest
	FolderID        string
	Domain          string
	CertificateName string
}

type YCCreatePostboxAddressRequest struct {
	YCBrowserRequest
	FolderID    string
	AddressName string
	PrivateKey  string
	Selector    string
	LogGroupID  string
}

type YCGetCertificateChallengeRequest struct {
	YCBrowserRequest
	CertificateID string
}

type YCBindCertificateToDNSRequest struct {
	YCBrowserRequest
	DNSZoneID        string
	DNSRecordName    string
	DNSRecordText    string
	DNSRecordOwnerID string
}

type YCGetPostboxVerificationRecordRequest struct {
	YCBrowserRequest
	IdentityID string
}

type YCBindApiGatewayToDNSRequest struct {
	YCBrowserRequest
	DNSZoneID        string
	DNSRecordName    string
	DNSRecordText    string
	DNSRecordOwnerID string
}

type YCCreateApiGatewayRequest struct {
	FolderID                 string
	Name                     string
	LogGroupID               string
	ProjectCapitalizeName    string
	FrontendStaticBucketName string
	ServiceAccountID         string
	BackendContainerID       string
}

type AddGithubRepositoryVariableRequest struct {
	GithubUserName  string
	GithubRepoName  string
	VariableName    string
	VariableValue   string
	EnvironmentName string
}

type AddGithubRepositoryEnvironmentRequest struct {
	GithubUserName  string
	GithubRepoName  string
	EnvironmentName string
}

type GetGithubRepositoryPublicKeyRequest struct {
	GithubUserName string
	GithubRepoName string
}

type AddGithubRepositorySecretRequest struct {
	GithubUserName         string
	GithubRepoName         string
	VariableName           string
	VariableEncryptedValue string
	RepositoryPublicKeyID  string
}

type GetGithubEnvironmentPublicKeyRequest struct {
	GithubUserName  string
	GithubRepoName  string
	EnvironmentName string
}

type AddGithubEnvironmentSecretRequest struct {
	GithubUserName         string
	GithubRepoName         string
	VariableName           string
	VariableEncryptedValue string
	RepositoryPublicKeyID  string
	EnvironmentName        string
}

type CreateSentryProjectRequest struct {
	ProjectName     string
	ProjectPlatform string
}

type CreateSentryProjectDSNRequest struct {
	ProjectName string
}

type UploadFileToBucketRequest struct {
	BucketName  string
	FileName    string
	FileContent []byte
	AccessKey   string
	SecretKey   string
}
