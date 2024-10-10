package model

type YCResponse struct {
	OperationID string
	Done        bool
}

type CreateFolderResponse struct {
	YCResponse
	FolderId string
}

type CreateBucketResponse struct {
	YCResponse
}

type CreateRegistryResponse struct {
	YCResponse
	RegistryID string
}

type CreateYDBResponse struct {
	YCResponse
	DatabaseEndpoint string
}

type CreateContainerResponse struct {
	YCResponse
	ContainerID string
}

type UpdateServiceAccountAccessToRegistryResponse struct {
	YCResponse
}

type CreateLoggingGroupResponse struct {
	YCResponse
	LogGroupID string
}

type CreateDNSZoneResponse struct {
	YCResponse
	DNSZoneId string
}

type CreateCertificateResponse struct {
	YCResponse
	CertificateID string
}

type CreateApiGatewayResponse struct {
	YCResponse
	ApiGatewayID string
}

type CreatePostboxAddressResponse struct {
	YCResponse
	PostboxAddressID string
	PostboxUsername  string
	PostboxPassword  string
}

type AddDNSRecordResponse struct {
	YCResponse
}

type AddDomainToGatewayResponse struct {
	YCResponse
}

type CreateAccessKeyResponse struct {
	AccessKeyID     string
	AccessKeySecret string
}
