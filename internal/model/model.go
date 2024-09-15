package model

import "fmt"

type Project struct {
	Name string
}

func (p *Project) getBackendRepoName() string {
	return fmt.Sprintf("%s-go", p.Name)
}

func (p *Project) getFrontendRepoName() string {
	return fmt.Sprintf("%s-vue3", p.Name)
}

func (p *Project) getBackendRepoUrl(githubDomain string, githubUserName string) string {
	return fmt.Sprintf("%s/%s/%s", githubDomain, githubUserName, p.getBackendRepoName())
}

func (p *Project) getFrontendRepoUrl(githubDomain string, githubUserName string) string {
	return fmt.Sprintf("%s/%s/%s", githubDomain, githubUserName, p.getFrontendRepoName())
}
