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

func (p *Service) SetupGithubFrontendRepo(ctx context.Context) error {
	err := functionswithneeds.Start(
		ctx,
		functionswithneeds.FunctionsWithNeeds{
			functionswithneeds.FunctionWithNeeds{
				Function: p.createGithubFrontendEnvironment,
				Needs:    nil,
			},
			functionswithneeds.FunctionWithNeeds{
				Function: p.createGithubFrontendRepositoryVariables,
				Needs:    nil,
			},
			functionswithneeds.FunctionWithNeeds{
				Function: p.createGithubFrontendRepositorySecrets,
				Needs:    nil,
			},
			functionswithneeds.FunctionWithNeeds{
				Function: p.createGithubFrontendEnvironmentVariables,
				Needs: []func(ctx context.Context) error{
					p.createGithubFrontendEnvironment,
				},
			},
			functionswithneeds.FunctionWithNeeds{
				Function: p.initAndPushLocalFrontendRepositoryToGithub,
				Needs: []func(ctx context.Context) error{
					p.createGithubFrontendRepositoryVariables,
					p.createGithubFrontendEnvironmentVariables,
				},
			},
		},
	)

	if err != nil {
		return err
	}

	p.logger.Info("Github: переменные для frontend репозитория заданы, репозиторий запушен")

	return nil
}

func (p *Service) createGithubFrontendEnvironment(ctx context.Context) error {
	return p.repositories.Github.AddRepositoryEnvironment(ctx, model.AddGithubRepositoryEnvironmentRequest{
		GithubUserName:  p.config.GitHub.UserName,
		GithubRepoName:  p.newProject.GetFrontendRepositoryName(),
		EnvironmentName: p.newProject.GetEnvironmentName(),
	})
}

func (p *Service) createGithubFrontendEnvironmentVariables(ctx context.Context) error {
	vars := model.FrontendEnvironmentVariables{
		Environment:  p.newProject.GetEnvironmentName(),
		APIURL:       p.newProject.GetDomainURL(),
		S3BucketName: p.newProject.GetYCObjectStorageBucketNameStatic(),
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
				GithubUserName: p.config.GitHub.UserName,
				GithubRepoName: p.newProject.GetFrontendRepositoryName(),
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

func (p *Service) createGithubFrontendRepositoryVariables(ctx context.Context) error {
	vars := model.FrontendRepositoryVariables{
		SentryDSN: "-",
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
				GithubRepoName: p.newProject.GetFrontendRepositoryName(),
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

func (p *Service) initAndPushLocalFrontendRepositoryToGithub(_ context.Context) error {
	dir, err := os.Getwd()

	if err != nil {
		return errors.WithStack(err)
	}

	_, err = exec.Command(
		"/bin/sh",
		fmt.Sprintf("%s/scripts/git-init-push.sh", dir),
		p.newProject.GetFrontendLocalPath(),
		p.newProject.GetFrontendGithubOrigin(),
	).Output()

	if err != nil {
		return errors.WithStack(err)
	}

	return nil
}

func (p *Service) createGithubFrontendRepositorySecrets(ctx context.Context) error {
	publicKey, err := p.repositories.Github.GetRepositoryPublicKey(ctx, model.GetGithubRepositoryPublicKeyRequest{
		GithubUserName: p.config.GitHub.UserName,
		GithubRepoName: p.newProject.GetFrontendRepositoryName(),
	})

	if err != nil {
		return err
	}

	vars := model.FrontendRepositorySecrets{
		NotificationsTelegramTo:    p.config.ProxyVariables.NotificationsTelegramTo,
		NotificationsTelegramToken: p.config.ProxyVariables.NotificationsTelegramToken,
		UpikothPackagesRead:        p.config.ProxyVariables.UpikothPackagesRead,
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
				GithubRepoName:         p.newProject.GetFrontendRepositoryName(),
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
