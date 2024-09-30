package model

type CreateFolderResponse struct {
	OperationID string
	FolderId    string
	Done        bool
}

type CreateBucketResponse struct {
	OperationID string
	Done        bool
}

type CreateRegistryResponse struct {
	OperationID string
	RegistryID  string
	Done        bool
}

type CreateYDBResponse struct {
	OperationID      string
	DatabaseEndpoint string
	Done             bool
}

type CreateContainerResponse struct {
	OperationID string
	ContainerID string
	Done        bool
}

type CreateLoggingGroupResponse struct {
	OperationID string
	LogGroupID  string
	Done        bool
}

type CreateDNSZoneResponse struct {
	OperationID string
	DNSZoneId   string
	Done        bool
}

type CreateCertificateResponse struct {
	OperationID   string
	CertificateID string
	Done          bool
}

type CreateApiGatewayResponse struct {
	OperationID string
	Done        bool
}

type CreatePostboxAddressResponse struct {
	OperationID string
	Done        bool
}
