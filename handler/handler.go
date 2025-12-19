package handler

import "codeedgeapp/config"

var conf config.Config

func Init(conf_ config.Config) {
	conf = conf_
}
