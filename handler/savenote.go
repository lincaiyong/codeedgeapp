package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/lincaiyong/larkbase"
	"github.com/lincaiyong/log"
	"os"
	"strconv"
)

func SaveNote(c *gin.Context) {
	var data struct {
		Note string `json:"note"`
		Id   string `json:"id"`
		Data string `json:"data"`
	}
	err := c.BindJSON(&data)
	if err != nil {
		errorResponse(c, "fail to bind json: %v", err)
		return
	}
	dataUrl, ok := conf.DataUrl[data.Data]
	if !ok {
		errorResponse(c, "data is invalid: %s", data.Data)
		return
	}
	conn, err := larkbase.ConnectAny(c.Request.Context(), os.Getenv("LARK_APP_ID"), os.Getenv("LARK_APP_SECRET"), dataUrl)
	if err != nil {
		errorResponse(c, "fail to connect: %v", err)
		return
	}
	id, _ := strconv.Atoi(data.Id)
	var record larkbase.AnyRecord
	err = conn.Find(&record, larkbase.NewFindOption(conn.FilterAnd(conn.Condition().Id.Is(id))))
	if err != nil {
		errorResponse(c, "fail to find record: %v", err)
		return
	}
	record.Update("note", data.Note)
	err = conn.Update(&record)
	if err != nil {
		errorResponse(c, "fail to update record: %v", err)
		return
	}
	log.InfoLog("record updated: %s", record.Id.StringValue())
	dataResponse(c, "done")
}
