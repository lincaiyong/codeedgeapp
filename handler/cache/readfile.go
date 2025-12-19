package cache

import (
	"os"
	"path/filepath"
)

func ReadFile(project, filePath string) ([]byte, error) {
	dir, err := EnsureProjectDir(project)
	if err != nil {
		return nil, err
	}
	filePath = filepath.Join(dir, filePath)
	b, err := os.ReadFile(filePath)
	return b, err
}
