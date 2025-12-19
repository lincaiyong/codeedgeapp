package cache

import (
	"os"
	"time"
)

func GetModTime(project string) (time.Time, error) {
	dir, err := EnsureProjectDir(project)
	if err != nil {
		return time.Time{}, err
	}
	stat, err := os.Stat(dir)
	if err != nil {
		return time.Time{}, err
	}
	return stat.ModTime(), nil
}
