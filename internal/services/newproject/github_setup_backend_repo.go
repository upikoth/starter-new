package newproject

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/upikoth/starter-new/internal/model"
	"golang.org/x/sync/errgroup"
	"os"
	"os/exec"
	"sync"
)

func (p *Service) SetupGithubBackendRepo(ctx context.Context) error {
	eg, newCtx := errgroup.WithContext(ctx)
	// 1. Устанавливаем env переменные в репозитории.
	eg.Go(func() error {
		return p.createBackendRepoVariables(newCtx)
	})

	// 2. Создаем environment в репозитории.
	eg.Go(func() error {
		return p.repositories.Github.AddRepositoryEnvironment(newCtx, model.AddGithubRepositoryEnvironmentRequest{
			GithubUserName:  p.config.GitHub.UserName,
			GithubRepoName:  p.newProject.GetBackendRepositoryName(),
			EnvironmentName: p.newProject.GetEnvironmentName(),
		})
	})

	err := eg.Wait()

	if err != nil {
		return err
	}

	// 3. Устанавливаем env переменные в репозитории для prod окружения.
	if err := p.createBackendEnvironmentVariables(ctx); err != nil {
		return err
	}

	// 4. Инициализируем git, пушим изменения в репозиторий.
	if err := p.initAndPushGitBackend(ctx); err != nil {
		return err
	}

	p.logger.Info("Github backend repository установлен")

	return nil
}

func (p *Service) createBackendEnvironmentVariables(ctx context.Context) error {
	environmentVariables := model.BackendRepositoryEnvironmentVariables{
		Environment: p.newProject.GetEnvironmentName(),
		FrontConfirmationPasswordRecoveryRequestURL: fmt.Sprintf("%s/#/auth/recovery-password-confirm", p.newProject.GetDomainURL()),
		FrontConfirmationRegistrationURL:            fmt.Sprintf("%s/#/auth/sign-up-confirm", p.newProject.GetDomainURL()),
		FrontURL:                                    p.newProject.GetDomain(),
		Port:                                        "8888",
		YCPFromAddress:                              p.newProject.GetEmailFromAddress(),
		YCPFromName:                                 p.newProject.GetEmailFromName(),
		YCPHost:                                     "postbox.cloud.yandex.net",
		YCPPort:                                     "25",
		YCS3Path:                                    p.newProject.GetYCObjectStorageBucketNameSecrets(),
		YDBAuthFileDirName:                          "/secrets",
		YDBAuthFile:                                 "authorized_key.json",
	}

	wg := sync.WaitGroup{}
	errs := make([]error, 0)

	bytes, err := json.Marshal(environmentVariables)
	if err != nil {
		return err
	}

	repoEnvironmentVariablesMap := map[string]string{}

	err = json.Unmarshal(bytes, &repoEnvironmentVariablesMap)
	if err != nil {
		return err
	}

	for k, v := range repoEnvironmentVariablesMap {
		wg.Add(1)
		go func() {
			err := p.repositories.Github.AddEnvironmentVariable(ctx, model.AddGithubRepositoryVariableRequest{
				GithubUserName: p.config.GitHub.UserName,
				GithubRepoName: p.newProject.GetBackendRepositoryName(),
				VariableName:   k,
				VariableValue:  v,
			})
			if err != nil {
				errs = append(errs, err)
			}
			defer wg.Done()
		}()
	}

	wg.Wait()
	if len(errs) > 0 {
		return errors.Join(errs...)
	}

	return nil
}

func (p *Service) createBackendRepoVariables(ctx context.Context) error {
	repoVariables := model.BackendRepositoryVariables{
		SentryDSN:              "-",
		YcContainterName:       p.newProject.GetYCServerlessContainerName(),
		YcFolderID:             p.newProject.GetYCFolderID(),
		YcLogOptionsLogGroupID: p.newProject.GetYCLoggingGroupID(),
		YcRegistry:             p.newProject.GetYCContainerRegistryID(),
		YcServiceAccountID:     p.newProject.GetYCServiceAccountID(),
	}
	wg := sync.WaitGroup{}
	errs := make([]error, 0)

	bytes, err := json.Marshal(repoVariables)
	if err != nil {
		return err
	}

	repoVariablesMap := map[string]string{}

	err = json.Unmarshal(bytes, &repoVariablesMap)
	if err != nil {
		return err
	}

	for k, v := range repoVariablesMap {
		wg.Add(1)
		go func() {
			err := p.repositories.Github.AddRepositoryVariable(ctx, model.AddGithubRepositoryVariableRequest{
				GithubUserName: p.config.GitHub.UserName,
				GithubRepoName: p.newProject.GetBackendRepositoryName(),
				VariableName:   k,
				VariableValue:  v,
			})
			if err != nil {
				errs = append(errs, err)
			}
			defer wg.Done()
		}()
	}

	wg.Wait()
	if len(errs) > 0 {
		return errors.Join(errs...)
	}

	return nil
}

func (p *Service) initAndPushGitBackend(_ context.Context) error {
	dir, err := os.Getwd()

	if err != nil {
		return err
	}

	_, err = exec.Command(
		"/bin/sh",
		fmt.Sprintf("%s/scripts/git-init-push.sh", dir),
		p.newProject.GetBackendLocalPath(),
		p.newProject.GetBackendGithubOrigin(),
	).Output()

	if err != nil {
		return err
	}

	return nil
}
