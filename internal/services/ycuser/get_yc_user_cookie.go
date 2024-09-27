package ycuser

import (
	"context"
	"fmt"
	"os"

	"github.com/upikoth/starter-new/internal/constants"
)

func (s *YCUserService) GetYcUserCookie(ctx context.Context) (string, error) {
	cookie := s.ycUser.getCookie()

	if cookie != "" {
		return cookie, nil
	}

	cookie, err := s.setYcUserCookie(ctx)

	if err != nil {
		return "", err
	}

	return cookie, nil
}

func (s *YCUserService) setYcUserCookie(ctx context.Context) (string, error) {
	path, err := os.Getwd()

	if err != nil {
		return "", err
	}

	cookie, err := s.repositories.FileInput.GetStringFromFile(ctx, fmt.Sprintf("%s/%s", path, constants.CookieFilename))

	if err != nil {
		return "", err
	}

	s.ycUser.setCookie(cookie)

	return cookie, nil
}
