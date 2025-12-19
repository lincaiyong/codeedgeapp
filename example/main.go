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
			"idor": "https://bytedance.larkoffice.com/base/P8QubLDkzabEJNsaNbacfha0nCd?table=tblUgfvHyAmuS3zx&view=vewDbnetVe",
		},
		DataFields: map[string][]string{
			"demo": {"sop"},
			"idor": {"sop", "version"},
		},
		SshRepoUrl: "git@code.byted.org:lincaiyong/samples",
		ChatFn:     monica.ChatCompletion,
		ResetCache: false,
	}
	monica.Init(os.Getenv("MONICA_SESSION_ID"))
	codeedgeapp.Run(conf)
}
