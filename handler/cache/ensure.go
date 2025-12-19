package cache

import (
	"context"
	"github.com/lincaiyong/daemon/common"
	"os"
	"path/filepath"
)

func ensureProjectDir(project string) (string, error) {
	dir := filepath.Join(cacheDir, project)
	if stat, err := os.Stat(dir); err == nil && stat.IsDir() {
		return dir, nil
	}
	err := os.MkdirAll(dir, os.ModePerm)
	if err != nil {
		return "'", err
	}
	_, _, err = common.RunCommand(context.Background(), dir, "git", "clone", "--depth=1", "--branch",
		project, sshRepoUrl, ".")
	if err != nil {
		return "", err
	}
	return dir, nil
}
