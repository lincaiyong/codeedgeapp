package cache

import "os"

const cacheDir = "/tmp/codeedgecache/"

var sshRepoUrl string

func Init(sshRepoUrl_ string, resetCache bool) {
	sshRepoUrl = sshRepoUrl_
	if resetCache {
		_ = os.RemoveAll(cacheDir)
	}
	_ = os.MkdirAll(cacheDir, os.ModePerm)
}
