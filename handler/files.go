package handler

import (
	"codeedgeapp/handler/cache"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/lincaiyong/gui"
	"net/http"
	"strings"
	"time"
)

func Files(c *gin.Context) {
	project := c.Query("project")
	vendor := c.Query("vendor")
	if project == "" || strings.Contains(project, ".") || strings.Contains(vendor, ".") {
		errorResponse(c, "project or vendor is invalid")
		return
	}
	projectMod, err := cache.GetModTime(project)
	if err != nil {
		errorResponse(c, "fail to get modified time: %v", err)
		return
	}
	if vendor == "" && gui.IfNotModifiedSince(c, projectMod) {
		c.String(http.StatusNotModified, "not modified")
		return
	}
	projects := []string{project}
	if vendor != "" {
		for _, item := range strings.Split(vendor, ",") {
			var vendorMod time.Time
			vendorMod, err = cache.GetModTime(item)
			if err != nil {
				errorResponse(c, "fail to get modified time: %v", err)
				return
			}
			if projectMod.Before(vendorMod) {
				projectMod = vendorMod
			}
			projects = append(projects, item)
		}
	}
	if gui.IfNotModifiedSince(c, projectMod) {
		c.String(http.StatusNotModified, "not modified")
		return
	}

	var result []string
	for _, name := range projects {
		var tmp []string
		tmp, err = cache.ReadFiles(name)
		if err != nil {
			errorResponse(c, "fail to read files: %v", err)
			return
		}
		if name != project {
			for _, item := range tmp {
				result = append(result, fmt.Sprintf("@vendor/%s/%s", name, item))
			}
		} else {
			result = append(result, tmp...)
		}
	}
	gui.SetLastModified(c, projectMod, 0)
	dataResponse(c, result)
}
