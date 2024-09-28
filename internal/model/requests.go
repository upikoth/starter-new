package model

type CreateCertificateRequest struct {
	FolderID        string
	Domain          string
	CertificateName string
	YCUserCookie    string
	YCUserCSRFToken string
}

type CreateApiGatewayRequest struct {
	FolderID                 string
	Name                     string
	LogGroupID               string
	ProjectCapitalizeName    string
	FrontendStaticBucketName string
	ServiceAccountID         string
	BackendContainerID       string
}
