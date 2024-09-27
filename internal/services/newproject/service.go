package newproject

import (
	"fmt"
	"strings"

	"github.com/upikoth/starter-new/internal/config"
	"github.com/upikoth/starter-new/internal/constants"
	"github.com/upikoth/starter-new/internal/pkg/logger"
	"github.com/upikoth/starter-new/internal/repositories"
	"github.com/upikoth/starter-new/internal/services/ycuser"
	"golang.org/x/text/cases"
	"golang.org/x/text/language"
)

type NewProjectService struct {
	newProject    *newProject
	logger        logger.Logger
	config        *config.Config
	repositories  *repositories.Repositories
	ycUserService *ycuser.YCUserService
}

func New(
	logger logger.Logger,
	config *config.Config,
	repositories *repositories.Repositories,
	ycUserService *ycuser.YCUserService,
) *NewProjectService {
	return &NewProjectService{
		logger:        logger,
		config:        config,
		repositories:  repositories,
		newProject:    &newProject{},
		ycUserService: ycUserService,
	}
}

func (p *NewProjectService) getBackendRepoName() string {
	return fmt.Sprintf("%s-go", p.newProject.name)
}

func (p *NewProjectService) getFrontendRepoName() string {
	return fmt.Sprintf("%s-vue3", p.newProject.name)
}

func (p *NewProjectService) getBackendRepoUrl() string {
	return fmt.Sprintf("%s/%s/%s", constants.GithubDomain, p.config.GitHub.UserName, p.getBackendRepoName())
}

func (p *NewProjectService) getFrontendRepoUrl() string {
	return fmt.Sprintf("%s/%s/%s", constants.GithubDomain, p.config.GitHub.UserName, p.getFrontendRepoName())
}

func (p *NewProjectService) getObjectStorageSecretsBucketName() string {
	return fmt.Sprintf("%s-secrets", p.newProject.name)
}

func (p *NewProjectService) getObjectStorageFrontendStaticBucketName() string {
	return p.getProjectSiteDomain()
}

func (p *NewProjectService) getProjectSiteDomain() string {
	return fmt.Sprintf("%s.%s", p.newProject.name, p.config.MainSiteDomainName)
}

func (p *NewProjectService) getProjectRegistryName() string {
	return p.newProject.name
}

func (p *NewProjectService) getProjectYDBName() string {
	return p.newProject.name
}

func (p *NewProjectService) getProjectServerlessContainerName() string {
	return p.newProject.name
}

func (p *NewProjectService) getProjectLoggingGroupName() string {
	return p.newProject.name
}

func (p *NewProjectService) getCertificateName() string {
	return strings.Join(strings.Split(p.getProjectSiteDomain(), "."), "-")
}

func (p *NewProjectService) getPostboxFromName() string {
	return cases.Title(language.English, cases.Compact).String(p.newProject.name)
}

func (p *NewProjectService) getPostboxFromAddress() string {
	return fmt.Sprintf("noreply@%s", p.getProjectSiteDomain())
}

func (p *NewProjectService) getProjectDNSZoneName() string {
	return fmt.Sprintf("%s.", p.getProjectSiteDomain())
}
