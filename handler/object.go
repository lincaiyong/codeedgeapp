package handler

import (
	_ "embed"
	"github.com/gin-gonic/gin"
	"github.com/lincaiyong/gui"
	"github.com/lincaiyong/log"
	"net/http"
	"time"
)

func GetObject(c *gin.Context) {
	key := c.Param("key")
	if gui.IfNotModifiedSince(c, time.Now().Add(-time.Hour*24*365)) {
		c.Status(http.StatusNotModified)
		return
	}
	data, err := conf.ObjectFn(c.Request.Context(), key)
	if err != nil {
		log.ErrorLog("fail to get object: %v", err)
		c.String(http.StatusNotFound, "object not found")
		return
	}
	c.Header("Content-Disposition", "attachment; filename="+key)
	c.Header("Content-Type", "application/octet-stream")
	gui.SetLastModified(c, time.Now(), 0)
	c.Data(http.StatusOK, "application/octet-stream", data)
}
