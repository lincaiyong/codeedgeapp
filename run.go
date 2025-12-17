package codeedgeapp

import (
	_ "embed"
	"github.com/gin-gonic/gin"
	"github.com/lincaiyong/daemon/common"
	. "github.com/lincaiyong/gui"
	"net/http"
	"os"
)

func Run(f func(group *gin.RouterGroup)) {
	common.StartServer(
		"codeedgeapp",
		"v1.0.1",
		"",
		func(envs []string, r *gin.RouterGroup) error {
			f(r)
			r.GET("/res/*filepath", HandleRes())
			r.GET("/", handlePage)
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
