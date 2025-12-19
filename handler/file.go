package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/lincaiyong/codeedgeapp/handler/cache"
	"github.com/lincaiyong/editdistance/edittool"
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
	patch := c.Query("patch")
	rhs := c.Query("rhs")
	if patch == "" {
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
	} else if rhs == "" {
		b, err := cache.ReadFile(project, filePath)
		if err != nil {
			errorResponse(c, "fail to read file: %v", err)
			return
		}
		oldStr := string(b)
		newStr := edittool.Patch(oldStr, patch)
		if oldStr == newStr {
			dataResponse(c, oldStr)
			return
		}
		dataResponse(c, [2]string{oldStr, newStr})
	} else {
		b, err := cache.ReadFile(project, filePath)
		if err != nil {
			errorResponse(c, "fail to read file: %v", err)
			return
		}
		oldStr := string(b)
		b, err = cache.ReadFile(rhs, filePath)
		if err != nil {
			errorResponse(c, "fail to read file: %v", err)
			return
		}
		newStr := string(b)
		if oldStr == newStr {
			dataResponse(c, oldStr)
			return
		}
		dataResponse(c, [2]string{oldStr, newStr})
	}
}
