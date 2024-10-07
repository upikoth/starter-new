package fileinput

import (
	"context"
	"github.com/pkg/errors"
	"os"
)

func (c *FileInput) GetStringFromFile(_ context.Context, filePath string) (string, error) {
	bytes, err := os.ReadFile(filePath)

	if err != nil {
		return "", errors.WithStack(err)
	}

	return string(bytes), nil
}
