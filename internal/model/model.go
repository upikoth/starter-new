package model

import "fmt"

type Project struct {
	Name             string
	FolderID         string
	ServiceAccountID string
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

func (p *Project) GetProjectSiteDomain(mainSiteDomainName string) string {
	return fmt.Sprintf("%s.%s", p.Name, mainSiteDomainName)
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
