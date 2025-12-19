package cache

import (
	"io/fs"
	"path/filepath"
)

func ReadFiles(project string) ([]string, error) {
	dir, err := ensureProjectDir(project)
	if err != nil {
		return nil, err
	}
	var result []string
	err = filepath.WalkDir(dir, func(path string, d fs.DirEntry, err error) error {
		if err != nil {
			return err
		}
		if d.IsDir() {
			return nil
		}
		result = append(result, path)
		return nil
	})
	return result, err
}
