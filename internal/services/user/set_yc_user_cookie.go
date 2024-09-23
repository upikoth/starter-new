package user

import (
	"context"
	"fmt"
	"os"

	"github.com/upikoth/starter-new/internal/model"
)

func (p *User) SetYcUserCookie(ctx context.Context) error {
	p.logger.Debug("Задаем cookie пользователя yandex cloud")

	path, err := os.Getwd()

	if err != nil {
		p.logger.Error(err.Error())
		return err
	}

	fmt.Println("Введите cookie пользователя yandex cloud. Их можно посмотреть через консоль разработчика в браузере")
	cookie, err := p.repositories.FileInput.GetStringFromFile(ctx, fmt.Sprintf("%s/cookie.txt", path))

	if err != nil {
		p.logger.Error("Не удалось задать cookie")
		return err
	}

	fmt.Println(cookie)

	p.YCUser.Cookie = model.Cookie(cookie)

	return nil
}
