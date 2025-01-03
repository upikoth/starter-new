package newproject

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/jefflinse/githubsecret"
	"github.com/pkg/errors"
	"github.com/upikoth/starter-new/internal/model"
	"github.com/upikoth/starter-new/internal/pkg/functionswithneeds"
	"os"
	"os/exec"
	"strings"
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
				Function: p.createGithubBackendRepositorySecrets,
				Needs:    nil,
			},
			functionswithneeds.FunctionWithNeeds{
				Function: p.createGithubBackendEnvironmentVariables,
				Needs: []func(ctx context.Context) error{
					p.createGithubBackendEnvironment,
				},
			},
			functionswithneeds.FunctionWithNeeds{
				Function: p.createGithubBackendEnvironmentSecrets,
				Needs: []func(ctx context.Context) error{
					p.createGithubBackendEnvironment,
				},
			},
			functionswithneeds.FunctionWithNeeds{
				Function: p.initAndPushLocalBackendRepositoryToGithub,
				Needs: []func(ctx context.Context) error{
					p.createGithubBackendRepositoryVariables,
					p.createGithubBackendEnvironmentVariables,
					p.createGithubBackendRepositorySecrets,
					p.createGithubBackendEnvironmentSecrets,
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
		FrontConfirmationPasswordRecoveryRequestURL: p.newProject.GetFrontendConfirmationPasswordRecoveryRequestURL(),
		FrontConfirmationRegistrationURL:            p.newProject.GetFrontendConfirmationRegistrationURL(),
		FrontURL:                                    p.newProject.GetDomain(),
		FrontHandleAuthPageURL:                      p.newProject.GetFrontendHandleAuthPageURL(),
		Port:                                        p.newProject.GetBackendPort(),
		YCPFromAddress:                              p.newProject.GetEmailFromAddress(),
		YCPFromName:                                 p.newProject.GetEmailFromName(),
		YCPHost:                                     p.newProject.GetYCPHost(),
		YCPPort:                                     p.newProject.GetYCPPort(),
		YCS3Path:                                    p.newProject.GetYCObjectStorageBucketNameSecrets(),
		YDBAuthFileDirName:                          p.newProject.GetYCYDBFileDirName(),
		YDBAuthFile:                                 p.newProject.GetYCYDBFileName(),
		OauthMailRedirectURL:                        "dummy",
		OauthVKRedirectURL:                          "dummy",
		OauthYandexRedirectURL:                      "dummy",
	}

	wg := sync.WaitGroup{}
	errs := make([]error, 0)

	bytes, err := json.Marshal(vars)
	if err != nil {
		return errors.WithStack(err)
	}

	varsMap := map[string]string{}

	err = json.Unmarshal(bytes, &varsMap)
	if err != nil {
		return errors.WithStack(err)
	}

	for k, v := range varsMap {
		wg.Add(1)
		go func() {
			err := p.repositories.Github.AddEnvironmentVariable(ctx, model.AddGithubRepositoryVariableRequest{
				GithubUserName:  p.config.GitHub.UserName,
				GithubRepoName:  p.newProject.GetBackendRepositoryName(),
				VariableName:    k,
				VariableValue:   v,
				EnvironmentName: p.newProject.GetEnvironmentName(),
			})
			if err != nil {
				errs = append(errs, err)
			}
			defer wg.Done()
		}()
	}

	wg.Wait()
	if len(errs) > 0 {
		errsString := make([]string, len(errs))
		for _, err := range errs {
			errsString = append(errsString, err.Error())
		}

		return errors.New(strings.Join(errsString, "\n"))
	}

	return nil
}

func (p *Service) createGithubBackendRepositoryVariables(ctx context.Context) error {
	vars := model.BackendRepositoryVariables{
		SentryDSN:              p.newProject.GetSentryBackendDSN(),
		YcContainterName:       p.newProject.GetYCServerlessContainerName(),
		YcFolderID:             p.newProject.GetYCFolderID(),
		YcLogOptionsLogGroupID: p.newProject.GetYCLoggingGroupID(),
		YcRegistry:             p.newProject.GetYCContainerRegistryID(),
		YcServiceAccountID:     p.newProject.GetYCServiceAccountID(),
		OauthMailAPIURL:        p.config.ProxyVariables.OauthMailAPIURL,
		OauthYandexAPIURL:      p.config.ProxyVariables.OauthYandexAPIURL,
	}
	wg := sync.WaitGroup{}
	errs := make([]error, 0)

	bytes, err := json.Marshal(vars)
	if err != nil {
		return errors.WithStack(err)
	}

	varsMap := map[string]string{}

	err = json.Unmarshal(bytes, &varsMap)
	if err != nil {
		return errors.WithStack(err)
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
		errsString := make([]string, len(errs))
		for _, err := range errs {
			errsString = append(errsString, err.Error())
		}

		return errors.New(strings.Join(errsString, "\n"))
	}

	return nil
}

func (p *Service) initAndPushLocalBackendRepositoryToGithub(_ context.Context) error {
	dir, err := os.Getwd()

	if err != nil {
		return errors.WithStack(err)
	}

	_, err = exec.Command(
		"/bin/sh",
		fmt.Sprintf("%s/scripts/git-init-push.sh", dir),
		p.newProject.GetBackendLocalPath(),
		p.newProject.GetBackendGithubOrigin(),
	).Output()

	if err != nil {
		return errors.WithStack(err)
	}

	return nil
}

func (p *Service) createGithubBackendRepositorySecrets(ctx context.Context) error {
	publicKey, err := p.repositories.Github.GetRepositoryPublicKey(ctx, model.GetGithubRepositoryPublicKeyRequest{
		GithubUserName: p.config.GitHub.UserName,
		GithubRepoName: p.newProject.GetBackendRepositoryName(),
	})

	if err != nil {
		return err
	}

	vars := model.BackendRepositorySecrets{
		NotificationsTelegramTo:    p.config.ProxyVariables.NotificationsTelegramTo,
		NotificationsTelegramToken: p.config.ProxyVariables.NotificationsTelegramToken,
		OauthMailClientID:          "dummy",
		OauthMailClientSecret:      "dummy",
		OauthVkClientID:            "dummy",
		OauthVkClientSecret:        "dummy",
		OauthYandexClientID:        "dummy",
		OauthYandexClientSecret:    "dummy",
	}
	wg := sync.WaitGroup{}
	errs := make([]error, 0)

	bytes, err := json.Marshal(vars)
	if err != nil {
		return errors.WithStack(err)
	}

	varsMap := map[string]string{}

	err = json.Unmarshal(bytes, &varsMap)
	if err != nil {
		return errors.WithStack(err)
	}

	for k, v := range varsMap {
		wg.Add(1)
		go func() {
			defer wg.Done()

			encryptedSecret, err := githubsecret.Encrypt(publicKey.Key, v)

			if err != nil {
				errs = append(errs, err)
				return
			}

			err = p.repositories.Github.AddRepositorySecret(ctx, model.AddGithubRepositorySecretRequest{
				GithubUserName:         p.config.GitHub.UserName,
				GithubRepoName:         p.newProject.GetBackendRepositoryName(),
				VariableName:           k,
				VariableEncryptedValue: encryptedSecret,
				RepositoryPublicKeyID:  publicKey.KeyID,
			})

			if err != nil {
				errs = append(errs, err)
			}
		}()
	}

	wg.Wait()
	if len(errs) > 0 {
		errsString := make([]string, len(errs))
		for _, err := range errs {
			errsString = append(errsString, err.Error())
		}

		return errors.New(strings.Join(errsString, "\n"))
	}

	return nil
}

func (p *Service) createGithubBackendEnvironmentSecrets(ctx context.Context) error {
	publicKey, err := p.repositories.Github.GetEnvironmentPublicKey(ctx, model.GetGithubEnvironmentPublicKeyRequest{
		GithubUserName:  p.config.GitHub.UserName,
		GithubRepoName:  p.newProject.GetBackendRepositoryName(),
		EnvironmentName: p.newProject.GetEnvironmentName(),
	})

	if err != nil {
		return err
	}

	vars := model.BackendEnvironmentSecrets{
		YCPostboxPassword:   p.newProject.GetYCPostboxPassword(),
		YCPostboxUsername:   p.newProject.GetYCPostboxUsername(),
		YCSAJSONCredentials: p.newProject.GetYCSAJSONCredentials(),
		YDBDSN:              p.newProject.GetYCYDBEndpoint(),
	}
	wg := sync.WaitGroup{}
	errs := make([]error, 0)

	bytes, err := json.Marshal(vars)
	if err != nil {
		return errors.WithStack(err)
	}

	varsMap := map[string]string{}

	err = json.Unmarshal(bytes, &varsMap)
	if err != nil {
		return errors.WithStack(err)
	}

	for k, v := range varsMap {
		wg.Add(1)
		go func() {
			defer wg.Done()

			encryptedSecret, err := githubsecret.Encrypt(publicKey.Key, v)

			if err != nil {
				errs = append(errs, err)
				return
			}

			err = p.repositories.Github.AddEnvironmentSecret(ctx, model.AddGithubEnvironmentSecretRequest{
				GithubUserName:         p.config.GitHub.UserName,
				GithubRepoName:         p.newProject.GetBackendRepositoryName(),
				VariableName:           k,
				VariableEncryptedValue: encryptedSecret,
				RepositoryPublicKeyID:  publicKey.KeyID,
				EnvironmentName:        p.newProject.GetEnvironmentName(),
			})

			if err != nil {
				errs = append(errs, err)
			}
		}()
	}

	wg.Wait()
	if len(errs) > 0 {
		errsString := make([]string, len(errs))
		for _, err := range errs {
			errsString = append(errsString, err.Error())
		}

		return errors.New(strings.Join(errsString, "\n"))
	}

	return nil
}
