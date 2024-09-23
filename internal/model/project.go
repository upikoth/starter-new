package model

import (
	"fmt"
	"strings"

	"golang.org/x/text/cases"
	"golang.org/x/text/language"
)

type Project struct {
	Name             string
	FolderID         string
	ServiceAccountID string
	RegistryID       string
	DatabaseEndpoint string
	LoggingGroupID   string
	PostboxUsername  string
	PostboxPassword  string
}

func (p *Project) GetBackendRepoName() string {
	return fmt.Sprintf("%s-go", p.Name)
}

func (p *Project) GetFrontendRepoName() string {
	return fmt.Sprintf("%s-vue3", p.Name)
}

func (p *Project) GetBackendRepoUrl(githubDomain string, githubUserName string) string {
	return fmt.Sprintf("%s/%s/%s", githubDomain, githubUserName, p.GetBackendRepoName())
}

func (p *Project) GetFrontendRepoUrl(githubDomain string, githubUserName string) string {
	return fmt.Sprintf("%s/%s/%s", githubDomain, githubUserName, p.GetFrontendRepoName())
}

func (p *Project) GetObjectStorageSecretsBucketName() string {
	return fmt.Sprintf("%s-secrets", p.Name)
}

func (p *Project) GetObjectStorageFrontendStaticBucketName(mainSiteDomainName string) string {
	return p.GetProjectSiteDomain(mainSiteDomainName)
}

func (p *Project) GetProjectSiteDomain(mainSiteDomainName string) string {
	return fmt.Sprintf("%s.%s", p.Name, mainSiteDomainName)
}

func (p *Project) GetProjectRegistryName() string {
	return p.Name
}

func (p *Project) GetProjectYDBName() string {
	return p.Name
}

func (p *Project) GetProjectServerlessContainerName() string {
	return p.Name
}

func (p *Project) GetProjectLoggingGroupName() string {
	return p.Name
}

func (p *Project) GetCertificateName(mainSiteDomainName string) string {
	return strings.Join(strings.Split(p.GetProjectSiteDomain(mainSiteDomainName), "."), "-")
}

func (p *Project) GetPostboxFromName() string {
	return cases.Title(language.English, cases.Compact).String(p.Name)
}

func (p *Project) GetPostboxFromAddress(mainSiteDomainName string) string {
	return fmt.Sprintf("noreply@%s", p.GetProjectSiteDomain(mainSiteDomainName))
}

func (p *Project) GetProjectDNSZoneName(mainSiteDomainName string) string {
	return fmt.Sprintf("%s.", p.GetProjectSiteDomain(mainSiteDomainName))
}
