package cache

import (
	"os"
)

const cacheDir = "/tmp/codeedgecache/"

var sshRepoUrl string

func Init(sshRepoUrl_ string) {
	_ = os.MkdirAll(cacheDir, os.ModePerm)
	sshRepoUrl = sshRepoUrl_
}
