package codeedgeapp

import (
	"codeedgeapp/chat"
	"codeedgeapp/config"
	"codeedgeapp/handler"
	"codeedgeapp/page"
	"context"
	_ "embed"
	"github.com/gin-gonic/gin"
	"github.com/lincaiyong/daemon/common"
	. "github.com/lincaiyong/gui"
	"net/http"
	"os"
)

func Run(conf config.Config) {
	handler.Init(conf)
	common.StartServer(
		"codeedgeapp",
		"v1.0.1",
		"",
		func(envs []string, r *gin.RouterGroup) error {
			r.GET("/res/*filepath", HandleRes())
			r.GET("/", page.Handle)
			r.GET("/files/", handler.Files)
			r.GET("/file/*filepath", handler.File)
			r.GET("/search/", handler.Search)
			r.POST("/chat/", handler.Chat)
			r.POST("/note/", handler.SaveNote)
			r.GET("/data/list", handler.ListData)
			r.GET("/data/:name", handler.Data)
			r.GET("/status", func(c *gin.Context) {
				c.JSON(http.StatusOK, gin.H{
					"status": "ok",
					"pid":    os.Getpid(),
				})
			})
			return nil
		},
	)
}
