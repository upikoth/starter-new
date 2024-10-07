package ycuser

import (
	"context"
	"github.com/pkg/errors"
	"net/url"
	"strings"
)

func (s *Service) GetYcUserCSRFToken(ctx context.Context) (string, error) {
	cookie, err := s.GetYcUserCookie(ctx)

	if err != nil {
		return "", errors.WithStack(err)
	}

	cookieMap := map[string]string{}

	for _, str := range strings.Split(cookie, ";") {
		keyValue := strings.Split(strings.Trim(str, " "), "=")

		cookieMap[strings.ToLower(keyValue[0])] = keyValue[1]
	}

	csrfToken, err := url.QueryUnescape(cookieMap["csrf-token"])

	if err != nil {
		return "", errors.WithStack(err)
	}

	return csrfToken, nil
}
