package newproject

import (
	"context"
)

func (p *NewProject) CreateYCFolder(ctx context.Context) error {
	p.logger.Info("Создаем новую папку проекта в yandex cloud")

	folderId, err := p.repositories.YandexCloud.CreateFolder(ctx, p.project.Name)

	if err != nil {
		p.logger.Error(err.Error())
		return err
	}

	p.project.FolderID = folderId
	p.logger.Info("Folder в yandex cloud успешно создан")
	return nil
}
