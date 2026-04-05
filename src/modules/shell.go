package modules

import (
	"os"
	"path/filepath"
)

func Shell() string {

	path := os.Getenv("SHELL")

	return filepath.Base(path)

}
