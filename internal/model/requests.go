package model

type CreateCertificateRequest struct {
	FolderID        string
	Domain          string
	CertificateName string
	YCUserCookie    string
	YCUserCSRFToken string
}
