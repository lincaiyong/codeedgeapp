package codeedgeapp

import (
	_ "embed"
	"github.com/gin-gonic/gin"
	"github.com/lincaiyong/codeedgeapp/handler"
	"github.com/lincaiyong/codeedgeapp/page"
	"github.com/lincaiyong/daemon/common"
	. "github.com/lincaiyong/gui"
)

func Run(conf handler.Config, admin bool) {
	handler.Init(conf)
	common.StartServer(
		"codeedgeapp",
		"v1.0.1",
		"",
		func(envs []string, r *gin.RouterGroup) error {
			r.GET("/res/*filepath", HandleRes())
			r.GET("/", page.Handle(admin))
			r.GET("/files/", handler.Files)
			r.GET("/diff/", handler.Diff)
			r.GET("/file/*filepath", handler.File)
			r.GET("/search/", handler.Search)
			r.POST("/chat/", handler.Chat)
			r.POST("/note/", handler.SaveNote)
			r.GET("/data/list/", handler.ListData)
			r.GET("/data/:name/", handler.Data)
			r.GET("/status/", handler.Status)
			return nil
		},
	)
}
