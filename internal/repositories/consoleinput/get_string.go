package consoleinput

import (
	"context"
	"fmt"
	"github.com/pkg/errors"
)

func (c *ConsoleInput) GetString(_ context.Context) (string, error) {
	str := ""

	if _, err := fmt.Scanln(&str); err != nil {
		return "", errors.WithStack(err)
	}

	return str, nil
}
