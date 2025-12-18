package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/lincaiyong/gui"
	"github.com/lincaiyong/log"
	"net/http"
	"path/filepath"
	"strings"
	"time"
)

func handleFiles(c *gin.Context) {
	project := c.Query("project")
	if project == "" || strings.Contains(project, ".") {
		c.String(http.StatusBadRequest, "project is invalid")
		return
	}

	filePath := filepath.Join("zip", project+".zip")
	mod, err := modifiedTime(filePath)
	if err != nil {
		log.ErrorLog("fail to get modified time: %v", err)
		c.String(http.StatusInternalServerError, "fail to stat zip")
		return
	}
	vendor := c.Query("vendor")
	if vendor == "" && gui.IfNotModifiedSince(c, mod) {
		c.String(http.StatusNotModified, "not modified")
		return
	}
	zipsToRead := []string{project}
	if vendor != "" {
		if strings.Contains(vendor, ".") {
			c.String(http.StatusBadRequest, "vendor is invalid")
			return
		}
		for _, item := range strings.Split(vendor, ",") {
			filePath = filepath.Join("zip", item+".zip")
			var itemMod time.Time
			itemMod, err = modifiedTime(filePath)
			if err != nil {
				log.ErrorLog("fail to get modified time: %v", err)
				c.String(http.StatusInternalServerError, "fail to stat zip")
				return
			}
			if mod.Before(itemMod) {
				mod = itemMod
			}
			zipsToRead = append(zipsToRead, item)
		}
	}
	if gui.IfNotModifiedSince(c, mod) {
		c.String(http.StatusNotModified, "not modified")
		return
	}

	ret := make([]string, 0)
	for _, k := range zipsToRead {
		path := filepath.Join("zip", k+".zip")
		var tmp []string
		tmp, err = readZipFiles(path)
		if err != nil {
			log.ErrorLog("fail to get sample: %v", err)
			c.String(http.StatusInternalServerError, "fail to read zip")
			return
		}
		if k != project {
			for _, item := range tmp {
				ret = append(ret, fmt.Sprintf("@vendor/%s/%s", strings.ReplaceAll(k, "/", "-"), item))
			}
		} else {
			ret = append(ret, tmp...)
		}
	}
	gui.SetLastModified(c, mod, 0)
	c.JSON(http.StatusOK, ret)
}

func handleFile(c *gin.Context) {
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
