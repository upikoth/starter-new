package ycuser

import (
	"context"
	"fmt"
	"github.com/pkg/errors"
	"os"

	"github.com/upikoth/starter-new/internal/constants"
)

func (s *Service) GetYcUserCookie(ctx context.Context) (string, error) {
	cookie := s.ycUser.GetCookie()

	if cookie != "" {
		return cookie, nil
	}

	cookie, err := s.setYcUserCookie(ctx)

	if err != nil {
		return "", errors.WithStack(err)
	}

	return cookie, nil
}

func (s *Service) setYcUserCookie(ctx context.Context) (string, error) {
	path, err := os.Getwd()

	if err != nil {
		return "", errors.WithStack(err)
	}

	cookie, err := s.repositories.FileInput.GetStringFromFile(ctx, fmt.Sprintf("%s/%s", path, constants.CookieFilename))

	if err != nil {
		return "", errors.WithStack(err)
	}

	s.ycUser.SetCookie(cookie)

	return cookie, nil
}
