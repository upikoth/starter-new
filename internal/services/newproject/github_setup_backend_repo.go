package newproject

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/upikoth/starter-new/internal/model"
	"github.com/upikoth/starter-new/internal/pkg/functionswithneeds"
	"os"
	"os/exec"
	"sync"
)

func (p *Service) SetupGithubBackendRepo(ctx context.Context) error {
	err := functionswithneeds.Start(
		ctx,
		functionswithneeds.FunctionsWithNeeds{
			functionswithneeds.FunctionWithNeeds{
				Function: p.createGithubBackendEnvironment,
				Needs:    nil,
			},
			functionswithneeds.FunctionWithNeeds{
				Function: p.createGithubBackendRepositoryVariables,
				Needs:    nil,
			},
			functionswithneeds.FunctionWithNeeds{
				Function: p.createGithubBackendEnvironmentVariables,
				Needs: []func(ctx context.Context) error{
					p.createGithubBackendEnvironment,
				},
			},
			functionswithneeds.FunctionWithNeeds{
				Function: p.initAndPushLocalBackendRepositoryToGithub,
				Needs: []func(ctx context.Context) error{
					p.createGithubBackendRepositoryVariables,
					p.createGithubBackendEnvironmentVariables,
				},
			},
		},
	)

	if err != nil {
		return err
	}

	p.logger.Info("Github: переменные для backend репозитория заданы, репозиторий запушен")

	return nil
}

func (p *Service) createGithubBackendEnvironment(ctx context.Context) error {
	return p.repositories.Github.AddRepositoryEnvironment(ctx, model.AddGithubRepositoryEnvironmentRequest{
		GithubUserName:  p.config.GitHub.UserName,
		GithubRepoName:  p.newProject.GetBackendRepositoryName(),
		EnvironmentName: p.newProject.GetEnvironmentName(),
	})
}

func (p *Service) createGithubBackendEnvironmentVariables(ctx context.Context) error {
	vars := model.BackendEnvironmentVariables{
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

	bytes, err := json.Marshal(vars)
	if err != nil {
		return err
	}

	varsMap := map[string]string{}

	err = json.Unmarshal(bytes, &varsMap)
	if err != nil {
		return err
	}

	for k, v := range varsMap {
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

func (p *Service) createGithubBackendRepositoryVariables(ctx context.Context) error {
	vars := model.BackendRepositoryVariables{
		SentryDSN:              "-",
		YcContainterName:       p.newProject.GetYCServerlessContainerName(),
		YcFolderID:             p.newProject.GetYCFolderID(),
		YcLogOptionsLogGroupID: p.newProject.GetYCLoggingGroupID(),
		YcRegistry:             p.newProject.GetYCContainerRegistryID(),
		YcServiceAccountID:     p.newProject.GetYCServiceAccountID(),
	}
	wg := sync.WaitGroup{}
	errs := make([]error, 0)

	bytes, err := json.Marshal(vars)
	if err != nil {
		return err
	}

	varsMap := map[string]string{}

	err = json.Unmarshal(bytes, &varsMap)
	if err != nil {
		return err
	}

	for k, v := range varsMap {
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

func (p *Service) initAndPushLocalBackendRepositoryToGithub(_ context.Context) error {
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
