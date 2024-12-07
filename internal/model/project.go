package model

import (
	"fmt"
	"github.com/upikoth/starter-new/internal/config"
	"golang.org/x/text/cases"
	"golang.org/x/text/language"
	"strings"
)

type Project struct {
	config *config.Config

	name                           string
	ycFolderID                     string
	ycServiceAccountID             string
	ycContainerRegistryID          string
	ydbEndpoint                    string
	ycLoggingGroupID               string
	ycCertificateID                string
	ycDNSZoneID                    string
	ycServerlessContainerID        string
	ycAPIGatewayID                 string
	ycPostboxAddressID             string
	ycPostboxUsername              string
	ycPostboxPassword              string
	ycObjectStorageAccessKeyID     string
	ycObjectStorageAccessKeySecret string
	ycSAJSONCredentials            string
	sentryBackendDSN               string
	sentryFrontendDSN              string
	githubBackendRepositoryID      int
	githubFrontendRepositoryID     int
}

func NewProject(config *config.Config) *Project {
	return &Project{
		config: config,
	}
}

const projectEnvironment = "prod"
const backendPort = "8888"
const ycpHost = "postbox.cloud.yandex.net"
const ycpPort = "25"
const ydbAuthFileDirName = "/secrets"
const ydbAuthFileName = "authorized_key.json"
const frontConfirmationPasswordRecoveryRequestAbsoluteURL = "/#/auth/recovery-password-confirm"
const frontConfirmationRegistrationAbsoluteURL = "/#/auth/sign-up-confirm"
const frontHandleAuthAbsoluteURL = "/#/auth/handle"
const swaggerAbsoluteURL = "/api/docs/app"

func (p *Project) GetName() string {
	return p.name
}

func (p *Project) SetName(name string) {
	p.name = name
}

func (p *Project) GetCapitalizeName() string {
	return cases.Title(language.English, cases.Compact).String(p.name)
}

func (p *Project) GetEnvironmentName() string {
	return projectEnvironment
}

func (p *Project) GetDomain() string {
	return fmt.Sprintf("%s.%s", p.GetName(), p.config.MainSiteDomain)
}

func (p *Project) GetDomainURL() string {
	return fmt.Sprintf("https://%s", p.GetDomain())
}

func (p *Project) GetLocalPath() string {
	return fmt.Sprintf("%s/%s", p.config.ProjectsLocalPath, p.GetName())
}

func (p *Project) GetEmailFromName() string {
	return cases.Title(language.English, cases.Compact).String(p.GetName())
}

func (p *Project) GetEmailFromAddress() string {
	return fmt.Sprintf("noreply@%s", p.GetDomain())
}

func (p *Project) GetBackendPort() string {
	return backendPort
}

func (p *Project) GetBackendRepositoryName() string {
	return fmt.Sprintf("%s-go", p.GetName())
}

func (p *Project) GetBackendLocalPath() string {
	return fmt.Sprintf("%s/%s/%s", p.config.ProjectsLocalPath, p.GetName(), p.GetBackendRepositoryName())
}

func (p *Project) GetBackendGithubOrigin() string {
	return fmt.Sprintf("git@github.com:%s/%s.git", p.config.GitHub.UserName, p.GetBackendRepositoryName())
}

func (p *Project) GetFrontendConfirmationPasswordRecoveryRequestURL() string {
	return fmt.Sprintf("%s%s", p.GetDomainURL(), frontConfirmationPasswordRecoveryRequestAbsoluteURL)
}

func (p *Project) GetFrontendConfirmationRegistrationURL() string {
	return fmt.Sprintf("%s%s", p.GetDomainURL(), frontConfirmationRegistrationAbsoluteURL)
}

func (p *Project) GetFrontendHandleAuthPageURL() string {
	return fmt.Sprintf("%s%s", p.GetDomainURL(), frontHandleAuthAbsoluteURL)
}

func (p *Project) GetSwaggerDocsURL() string {
	return fmt.Sprintf("%s%s", p.GetDomainURL(), swaggerAbsoluteURL)
}

func (p *Project) GetFrontendRepositoryName() string {
	return fmt.Sprintf("%s-vue3", p.GetName())
}

func (p *Project) GetFrontendLocalPath() string {
	return fmt.Sprintf("%s/%s/%s", p.config.ProjectsLocalPath, p.GetName(), p.GetFrontendRepositoryName())
}

func (p *Project) GetFrontendGithubOrigin() string {
	return fmt.Sprintf("git@github.com:%s/%s.git", p.config.GitHub.UserName, p.GetFrontendRepositoryName())
}

func (p *Project) GetYCPHost() string {
	return ycpHost
}

func (p *Project) GetYCPPort() string {
	return ycpPort
}

func (p *Project) GetYCYDBFileDirName() string {
	return ydbAuthFileDirName
}

func (p *Project) GetYCYDBFileName() string {
	return ydbAuthFileName
}

func (p *Project) GetYCObjectStorageBucketNameSecrets() string {
	return fmt.Sprintf("%s-secrets", p.GetName())
}

func (p *Project) GetYCObjectStorageBucketNameStatic() string {
	return p.GetDomain()
}

func (p *Project) GetYCContainerRegistryName() string {
	return p.GetName()
}

func (p *Project) GetYCYDBName() string {
	return p.GetName()
}

func (p *Project) GetYCServerlessContainerName() string {
	return p.GetName()
}

func (p *Project) GetYCLoggingGroupName() string {
	return p.GetName()
}

func (p *Project) GetYCCertificateName() string {
	return strings.Join(strings.Split(p.GetDomain(), "."), "-")
}

func (p *Project) GetYCPostboxName() string {
	return p.GetDomain()
}

func (p *Project) GetYCDNSZoneName() string {
	return fmt.Sprintf("%s.", p.GetDomain())
}

func (p *Project) GetYCApiGatewayName() string {
	return p.GetName()
}

func (p *Project) GetYCFolderID() string {
	return p.ycFolderID
}

func (p *Project) SetYCFolderID(id string) {
	p.ycFolderID = id
}

func (p *Project) GetYCServiceAccountID() string {
	return p.ycServiceAccountID
}

func (p *Project) SetYCServiceAccountID(id string) {
	p.ycServiceAccountID = id
}

func (p *Project) GetYCContainerRegistryID() string {
	return p.ycContainerRegistryID
}

func (p *Project) SetYCContainerRegistryID(id string) {
	p.ycContainerRegistryID = id
}

func (p *Project) GetYCLoggingGroupID() string {
	return p.ycLoggingGroupID
}

func (p *Project) SetYCLoggingGroupID(id string) {
	p.ycLoggingGroupID = id
}

func (p *Project) GetYCCertificateID() string {
	return p.ycCertificateID
}

func (p *Project) SetYCCertificateID(id string) {
	p.ycCertificateID = id
}

func (p *Project) GetYCDNSZoneID() string {
	return p.ycDNSZoneID
}

func (p *Project) SetYCDNSZoneID(id string) {
	p.ycDNSZoneID = id
}

func (p *Project) GetYCServerlessContainerID() string {
	return p.ycServerlessContainerID
}

func (p *Project) SetYCServerlessContainerID(id string) {
	p.ycServerlessContainerID = id
}

func (p *Project) GetYCAPIGatewayID() string {
	return p.ycAPIGatewayID
}

func (p *Project) SetYCAPIGatewayID(id string) {
	p.ycAPIGatewayID = id
}

func (p *Project) GetYCYDBEndpoint() string {
	return p.ydbEndpoint
}

func (p *Project) SetYCYDBEndpoint(e string) {
	p.ydbEndpoint = e
}

func (p *Project) GetYCPostboxAddressID() string {
	return p.ycPostboxAddressID
}

func (p *Project) SetYCPostboxAddressID(id string) {
	p.ycPostboxAddressID = id
}

func (p *Project) GetYCPostboxUsername() string {
	return p.ycPostboxUsername
}

func (p *Project) SetYCPostboxUsername(username string) {
	p.ycPostboxUsername = username
}

func (p *Project) GetYCPostboxPassword() string {
	return p.ycPostboxPassword
}

func (p *Project) SetYCPostboxPassword(password string) {
	p.ycPostboxPassword = password
}

func (p *Project) GetYCObjectStorageAccessKeyID() string {
	return p.ycObjectStorageAccessKeyID
}

func (p *Project) SetYCObjectStorageAccessKeyID(id string) {
	p.ycObjectStorageAccessKeyID = id
}

func (p *Project) GetYCObjectStorageAccessKeySecret() string {
	return p.ycObjectStorageAccessKeySecret
}

func (p *Project) SetYCObjectStorageAccessKeySecret(secret string) {
	p.ycObjectStorageAccessKeySecret = secret
}

func (p *Project) GetYCSAJSONCredentials() string {
	return p.ycSAJSONCredentials
}

func (p *Project) SetYCSAJSONCredentials(creds string) {
	p.ycSAJSONCredentials = creds
}

func (p *Project) GetSentryBackendDSN() string {
	return p.sentryBackendDSN
}

func (p *Project) SetSentryBackendDSN(dsn string) {
	p.sentryBackendDSN = dsn
}

func (p *Project) GetSentryFrontendDSN() string {
	return p.sentryFrontendDSN
}

func (p *Project) SetSentryFrontendDSN(dsn string) {
	p.sentryFrontendDSN = dsn
}

func (p *Project) GetGithubBackendRepositoryID() int {
	return p.githubBackendRepositoryID
}

func (p *Project) SetGithubBackendRepositoryID(id int) {
	p.githubBackendRepositoryID = id
}

func (p *Project) GetGithubFrontendRepositoryID() int {
	return p.githubFrontendRepositoryID
}

func (p *Project) SetGithubFrontendRepositoryID(id int) {
	p.githubFrontendRepositoryID = id
}
