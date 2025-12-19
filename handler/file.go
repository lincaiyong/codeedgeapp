package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/lincaiyong/codeedgeapp/handler/cache"
	"github.com/lincaiyong/gui"
	"net/http"
	"strings"
)

func File(c *gin.Context) {
	filePath := c.Param("filepath")[1:]
	project := c.Query("project")
	if project == "" || strings.Contains(project, ".") {
		errorResponse(c, "project is invalid: %s", project)
		return
	}
	mod, err := cache.GetModTime(project)
	if err != nil {
		errorResponse(c, "project not found")
		return
	}
	if gui.IfNotModifiedSince(c, mod) {
		c.String(http.StatusNotModified, "not modified")
		return
	}
	b, err := cache.ReadFile(project, filePath)
	if err != nil {
		errorResponse(c, "fail to read file: %v", err)
		return
	}
	gui.SetLastModified(c, mod, 0)
	dataResponse(c, string(b))
}
