package codeedgeapp

import (
	_ "embed"
	"github.com/gin-gonic/gin"
	"github.com/lincaiyong/daemon/common"
	. "github.com/lincaiyong/gui"
)

func Run() {
	common.StartServer(
		"codeedgeapp",
		"v1.0.1",
		"",
		func(envs []string, r *gin.RouterGroup) error {
			r.GET("/res/*filepath", HandleRes())
			r.GET("/", handlePage)
			return nil
		},
	)
}
