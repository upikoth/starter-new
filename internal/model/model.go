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
