package handler

import (
	"codeedgeapp/handler/cache"
	"github.com/gin-gonic/gin"
	"github.com/lincaiyong/gui"
	"net/http"
	"strings"
)

func File(c *gin.Context) {
	var project string
	filePath := c.Param("filepath")
	pathItems := strings.Split(filePath, "/")
	if pathItems[1] == "@vendor" {
		if len(pathItems) > 3 {
			project = pathItems[2]
			filePath = strings.Join(pathItems[3:], "/")
		} else {
			errorResponse(c, "file not found")
			return
		}
	} else {
		project = c.Query("project")
		filePath = strings.Join(pathItems[1:], "/")
	}
	if project == "" || strings.Contains(project, ".") {
		errorResponse(c, "project is invalid: %s", project)
		return
	}
	mod, err := cache.GetModTime(project)
	if err != nil {
		errorResponse(c, "fail to get file modtime: %v", err)
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
	c.String(http.StatusOK, string(b))
}
