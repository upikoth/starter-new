package model

import "fmt"

type Project struct {
	Name             string
	FolderID         string
	ServiceAccountID string
	RegistryID       string
	DatabaseEndpoint string
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
	Done        bool
}
