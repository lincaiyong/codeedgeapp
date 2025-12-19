package handler

import (
	"codeedgeapp/config"
	"codeedgeapp/handler/cache"
	"fmt"
	"strings"
)

var conf config.Config

func Init(conf_ config.Config) {
	sshRepoUrl := fmt.Sprintf("git@%s", strings.Replace(conf_.SamplesRepo, "/", ":", 1))
	cache.Init(sshRepoUrl)
	conf = conf_
}
