package fileinput

import (
	"context"
	"os"
)

func (c *FileInput) GetStringFromFile(_ context.Context, filePath string) (string, error) {
	bytes, err := os.ReadFile(filePath)

	if err != nil {
		c.logger.Error(err.Error())
		return "", err
	}

	return string(bytes), nil
}
