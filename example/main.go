package main

import (
	"github.com/lincaiyong/codeedgeapp"
	"github.com/lincaiyong/codeedgeapp/config"
	"github.com/lincaiyong/uniapi/service/monica"
	"os"
)

func main() {
	conf := config.Config{
		AppId:     os.Getenv("LARK_APP_ID"),
		AppSecret: os.Getenv("LARK_APP_ID"),
		DataUrl: map[string]string{
			"demo": "https://bytedance.larkoffice.com/base/RB31bsA7Pa3f5JsKDlhcoTYdnue?table=tblxbNmiqJl67Egt&view=vewQotpDmR",
		},
		DataFields: map[string][]string{},
		SshRepoUrl: "git@github.com:lincaiyong/samples",
		ChatFn:     monica.ChatCompletion,
	}
	monica.Init(os.Getenv("MONICA_SESSION_ID"))
	codeedgeapp.Run(conf)
}
