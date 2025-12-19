package main

import (
	"github.com/lincaiyong/codeedgeapp"
	"github.com/lincaiyong/codeedgeapp/handler"
	"github.com/lincaiyong/uniapi/service/monica"
	"os"
)

func main() {
	conf := handler.Config{
		AppId:     os.Getenv("LARK_APP_ID"),
		AppSecret: os.Getenv("LARK_APP_SECRET"),
		DataUrl: map[string]string{
			"demo": "https://bytedance.larkoffice.com/base/RB31bsA7Pa3f5JsKDlhcoTYdnue?table=tblxbNmiqJl67Egt&view=vewQotpDmR",
		},
		DataFields: map[string][]string{},
		SshRepoUrl: "git@github.com:lincaiyong/samples",
		ChatFn:     monica.ChatCompletion,
		ResetCache: false,
	}
	monica.Init(os.Getenv("MONICA_SESSION_ID"))
	codeedgeapp.Run(conf)
}
