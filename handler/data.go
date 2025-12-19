package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/lincaiyong/larkbase"
)

var requiredFields = []string{"project", "vendor", "note"}

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
	result := make([]map[string]string, 0)
	fields := append(requiredFields, conf.DataFields[name]...)
	for _, record := range records {
		item := map[string]string{
			"id": record.Id.StringValue(),
		}
		for _, field := range fields {
			item[field] = record.Data[field]
		}
		result = append(result, item)
	}
	dataResponse(c, result)
}
