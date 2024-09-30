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

type CreatePostboxAddressRequest struct {
	FolderID        string
	AddressName     string
	YCUserCookie    string
	YCUserCSRFToken string
	PrivateKey      string
	Selector        string
	LogGroupID      string
}

type GetCertificateChallengeRequest struct {
	CertificateID   string
	YCUserCookie    string
	YCUserCSRFToken string
}

type BindCertificateToDNSRequest struct {
	DNSZoneID        string
	DNSRecordName    string
	DNSRecordText    string
	DNSRecordOwnerID string
	YCUserCookie     string
	YCUserCSRFToken  string
}
