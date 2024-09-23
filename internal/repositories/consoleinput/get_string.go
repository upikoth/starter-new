package consoleinput

import (
	"context"
	"fmt"
)

func (c *ConsoleInput) GetString(_ context.Context) (string, error) {
	str := ""

	if _, err := fmt.Scanln(&str); err != nil {
		c.logger.Error(err.Error())
		return "", err
	}

	return str, nil
}
