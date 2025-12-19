package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/lincaiyong/gui"
	"github.com/lincaiyong/log"
	"net/http"
	"path/filepath"
	"strings"
)

func File(c *gin.Context) {
	var project string
	filePath := c.Param("filepath")[1:]
	if strings.HasPrefix(filePath, "@vendor/") {
		filePath = strings.Replace(filePath, "@vendor/", "", 1)
		if idx := strings.Index(filePath, "/"); idx == -1 {
			c.String(http.StatusNotFound, "file not found")
			return
		} else {
			project = strings.ReplaceAll(filePath[:idx], "-", "/")
			filePath = filePath[idx+1:]
		}
	} else {
		project = c.Query("project")
	}
	if project == "" || strings.Contains(project, ".") {
		c.String(http.StatusBadRequest, "project is invalid")
		return
	}
	zipFilePath := filepath.Join("zip", project+".zip")
	mod, err := modifiedTime(zipFilePath)
	if err != nil {
		log.ErrorLog("fail to get modified time: %v", err)
		c.String(http.StatusInternalServerError, "fail to stat zip")
		return
	}
	if gui.IfNotModifiedSince(c, mod) {
		c.String(http.StatusNotModified, "not modified")
		return
	}
	b, err := readZipFile(zipFilePath, filePath)
	if err != nil {
		log.ErrorLog("fail to read file: %v", err)
		c.String(http.StatusInternalServerError, "fail to read file")
		return
	}
	gui.SetLastModified(c, mod, 0)
	c.String(http.StatusOK, string(b))
}
