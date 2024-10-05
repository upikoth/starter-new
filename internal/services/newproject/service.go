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

type Service struct {
	newProject    *newProject
	logger        logger.Logger
	config        *config.Config
	repositories  *repositories.Repositories
	ycUserService *ycuser.Service
}

func New(
	logger logger.Logger,
	config *config.Config,
	repositories *repositories.Repositories,
	ycUserService *ycuser.Service,
) *Service {
	return &Service{
		logger:        logger,
		config:        config,
		repositories:  repositories,
		newProject:    &newProject{},
		ycUserService: ycUserService,
	}
}

func (p *Service) getBackendRepoName() string {
	return fmt.Sprintf("%s-go", p.newProject.name)
}

func (p *Service) getFrontendRepoName() string {
	return fmt.Sprintf("%s-vue3", p.newProject.name)
}

func (p *Service) getBackendRepoUrl() string {
	return fmt.Sprintf("%s/%s/%s", constants.GithubDomain, p.config.GitHub.UserName, p.getBackendRepoName())
}

func (p *Service) getFrontendRepoUrl() string {
	return fmt.Sprintf("%s/%s/%s", constants.GithubDomain, p.config.GitHub.UserName, p.getFrontendRepoName())
}

func (p *Service) getObjectStorageSecretsBucketName() string {
	return fmt.Sprintf("%s-secrets", p.newProject.name)
}

func (p *Service) getObjectStorageFrontendStaticBucketName() string {
	return p.getProjectSiteDomain()
}

func (p *Service) getProjectSiteDomain() string {
	return fmt.Sprintf("%s.%s", p.newProject.name, p.config.MainSiteDomainName)
}

func (p *Service) getProjectRegistryName() string {
	return p.newProject.name
}

func (p *Service) getProjectYDBName() string {
	return p.newProject.name
}

func (p *Service) getProjectServerlessContainerName() string {
	return p.newProject.name
}

func (p *Service) getProjectLoggingGroupName() string {
	return p.newProject.name
}

func (p *Service) getCertificateName() string {
	return strings.Join(strings.Split(p.getProjectSiteDomain(), "."), "-")
}

func (p *Service) getPostboxFromName() string {
	return cases.Title(language.English, cases.Compact).String(p.newProject.name)
}

func (p *Service) getPostboxFromAddress() string {
	return fmt.Sprintf("noreply@%s", p.getProjectSiteDomain())
}

func (p *Service) getPostboxAddressName() string {
	return p.getProjectSiteDomain()
}

func (p *Service) getProjectDNSZoneName() string {
	return fmt.Sprintf("%s.", p.getProjectSiteDomain())
}

func (p *Service) getApiGatewayName() string {
	return p.newProject.name
}

func (p *Service) getCapitalizeName() string {
	return cases.Title(language.English, cases.Compact).String(p.newProject.name)
}

func (p *Service) getProjectLocalPath() string {
	return fmt.Sprintf("%s/%s", p.config.ProjectsPathLocal, p.newProject.name)
}

func (p *Service) getProjectLocalPathBackend() string {
	return fmt.Sprintf("%s/%s/%s", p.config.ProjectsPathLocal, p.newProject.name, p.getBackendRepoName())
}

func (p *Service) getProjectLocalPathFrontend() string {
	return fmt.Sprintf("%s/%s/%s", p.config.ProjectsPathLocal, p.newProject.name, p.getFrontendRepoName())
}

func (p *Service) getProjectGithubOriginBackend() string {
	return fmt.Sprintf("git@github.com:%s/%s.git", p.config.GitHub.UserName, p.getBackendRepoName())
}

func (p *Service) getProjectGithubOriginFrontend() string {
	return fmt.Sprintf("git@github.com:%s/%s.git", p.config.GitHub.UserName, p.getFrontendRepoName())
}
