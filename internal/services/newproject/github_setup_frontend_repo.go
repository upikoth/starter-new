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

func (p *Service) SetupGithubFrontendRepo(ctx context.Context) error {
	eg, newCtx := errgroup.WithContext(ctx)
	// 1. Устанавливаем env переменные в репозитории.
	eg.Go(func() error {
		return p.createFrontendRepoVariables(newCtx)
	})

	// 2. Создаем environment в репозитории.
	eg.Go(func() error {
		return p.repositories.Github.AddRepositoryEnvironment(newCtx, model.AddGithubRepositoryEnvironmentRequest{
			GithubUserName:  p.config.GitHub.UserName,
			GithubRepoName:  p.newProject.GetFrontendRepositoryName(),
			EnvironmentName: p.newProject.GetEnvironmentName(),
		})
	})

	err := eg.Wait()

	if err != nil {
		return err
	}

	// 3. Устанавливаем env переменные в репозитории для prod окружения.
	if err := p.createFrontendEnvironmentVariables(ctx); err != nil {
		return err
	}

	// 4. Инициализируем git, пушим изменения в репозиторий.
	if err := p.initAndPushGitFrontend(ctx); err != nil {
		return err
	}

	p.logger.Info("Github frontend repository установлен")

	return nil
}

func (p *Service) createFrontendEnvironmentVariables(ctx context.Context) error {
	environmentVariables := model.FrontendRepositoryEnvironmentVariables{
		Environment:  p.newProject.GetEnvironmentName(),
		APIURL:       p.newProject.GetDomainURL(),
		S3BucketName: p.newProject.GetYCObjectStorageBucketNameStatic(),
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
		return errors.Join(errs...)
	}

	return nil
}

func (p *Service) createFrontendRepoVariables(ctx context.Context) error {
	repoVariables := model.FrontendRepositoryVariables{
		SentryDSN: "-",
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
		return errors.Join(errs...)
	}

	return nil
}

func (p *Service) initAndPushGitFrontend(_ context.Context) error {
	dir, err := os.Getwd()

	if err != nil {
		return err
	}

	_, err = exec.Command(
		"/bin/sh",
		fmt.Sprintf("%s/scripts/git-init-push.sh", dir),
		p.newProject.GetFrontendLocalPath(),
		p.newProject.GetFrontendGithubOrigin(),
	).Output()

	if err != nil {
		return err
	}

	return nil
}
