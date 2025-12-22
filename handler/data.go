package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/lincaiyong/gui"
	"github.com/lincaiyong/larkbase"
	"net/http"
	"time"
)

var requiredFields = []string{"id", "project", "vendor", "note"}

func Data(c *gin.Context) {
	name := c.Param("name")
	url := conf.DataUrl[name]
	if url == "" {
		errorResponse(c, "data not found: %s", name)
		return
	}
	conn, err := larkbase.ConnectAny(c.Request.Context(), conf.AppId, conf.AppSecret, url)
	if err != nil {
		errorResponse(c, "fail to connect: %v", err)
		return
	}
	var records []*larkbase.AnyRecord
	err = conn.FindAll(&records, nil)
	if err != nil {
		errorResponse(c, "fail to query: %v", err)
		return
	}
	result := make([][]string, 0)
	fields := append(requiredFields, conf.DataFields[name]...)
	var mod time.Time
	for _, record := range records {
		row := make([]string, 0)
		for _, field := range fields {
			row = append(row, record.Data[field])
		}
		result = append(result, row)
		if mod.Before(record.ModifiedTime) {
			mod = record.ModifiedTime
		}
	}
	if gui.IfNotModifiedSince(c, mod) {
		c.Status(http.StatusNotModified)
		return
	}
	gui.SetLastModified(c, mod, 0)
	dataResponse(c, gin.H{
		"fields": fields,
		"data":   result,
	})
}
