package main

import (
	_ "embed"
	"github.com/gin-gonic/gin"
	"github.com/lincaiyong/daemon/common"
)

func main() {
	common.StartServer(
		"sample",
		"v1.0.1",
		"",
		func(envs []string, r *gin.RouterGroup) error {
			r.GET("/files/", handleFiles)
			r.GET("/file/*filepath", handleFile)
			r.GET("/search/", handleSearch)
			return nil
		},
	)
}
