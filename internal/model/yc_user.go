package model

import (
	"net/url"
	"strings"
)

type Cookie string

type YCUser struct {
	Cookie Cookie
}

func (c Cookie) GetValue() string {
	return string(c)
}

func (c Cookie) GetCSRFToken() string {
	cookieMap := map[string]string{}

	for _, str := range strings.Split(c.GetValue(), ";") {
		keyValue := strings.Split(strings.Trim(str, " "), "=")

		cookieMap[keyValue[0]] = keyValue[1]
	}

	res, _ := url.QueryUnescape(cookieMap["CSRF-TOKEN"])

	return res
}
