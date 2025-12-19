package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/lincaiyong/codeedgeapp/handler/cache"
	"github.com/lincaiyong/gui"
	"net/http"
	"strings"
)

func Files(c *gin.Context) {
	project := c.Query("project")
	if project == "" || strings.Contains(project, ".") {
		errorResponse(c, "project is invalid")
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
	result, err := cache.ReadFiles(project)
	if err != nil {
		errorResponse(c, "fail to read files: %v", err)
		return
	}
	gui.SetLastModified(c, mod, 0)
	dataResponse(c, result)
}
