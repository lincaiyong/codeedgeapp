package handler

import (
	"github.com/lincaiyong/codeedgeapp/config"
	"github.com/lincaiyong/codeedgeapp/handler/cache"
)

var conf config.Config

func Init(conf_ config.Config) {
	cache.Init(conf.SshRepoUrl)
	conf = conf_
}
