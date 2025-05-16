package asciify

import (
	"errors"
	"io/fs"
	"os"
)

func fileNotExists(path string) bool {
	_, err := os.Stat(path)
	return err != nil && errors.Is(err, fs.ErrNotExist)
}
