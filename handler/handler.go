package handler

import (
	"github.com/lincaiyong/codeedgeapp/handler/cache"
)

var conf Config

func Init(conf_ Config) {
	conf = conf_
	cache.Init(conf.SshRepoUrl, conf.ResetCache)
}
