package main

import "github.com/lincaiyong/codeedgeapp"

func main() {
	data := map[string]string{
		"demo": "https://bytedance.larkoffice.com/base/RB31bsA7Pa3f5JsKDlhcoTYdnue?table=tblxbNmiqJl67Egt&view=vewQotpDmR",
	}
	samplesRepo := "github.com/lincaiyong/samples"
	codeedgeapp.Run(data, samplesRepo)
}
