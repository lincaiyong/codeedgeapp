package main

import (
	"github.com/gin-gonic/gin"
	"github.com/lincaiyong/arg"
	"github.com/lincaiyong/codeedgeapp"
)

func main() {
	arg.Parse()
	download := arg.KeyValueArg("download", "")
	if download != "" {
		Download(download)
		return
	}
	codeedgeapp.Run(func(r *gin.RouterGroup) {
		r.GET("/files/", handleFiles)
		r.GET("/file/*filepath", handleFile)
		r.GET("/search/", handleSearch)
		r.POST("/chat/", handleChat)
		r.POST("/note/", handleSaveNote)
		r.GET("/data/", handleData)
	})
}
