package main

import (
	"github.com/gin-gonic/gin"
	"github.com/lincaiyong/larkbase"
	"github.com/lincaiyong/log"
	"net/http"
	"os"
	"strconv"
)

type Record struct {
	larkbase.Meta `lark:"https://bytedance.larkoffice.com/base/RB31bsA7Pa3f5JsKDlhcoTYdnue?table=tblxbNmiqJl67Egt"`
	Id            larkbase.NumberField       `lark:"id"`
	Sop           larkbase.TextField         `lark:"sop"`
	Note          larkbase.TextField         `lark:"note"`
	Project       larkbase.TextField         `lark:"project"`
	Vendor        larkbase.TextField         `lark:"vendor"`
	Version       larkbase.SingleSelectField `lark:"version"`
}

func handleSaveNote(c *gin.Context) {
	var data struct {
		Note string `json:"note"`
		Id   string `json:"id"`
	}
	err := c.BindJSON(&data)
	if err != nil {
		log.ErrorLog("fail to bind json: %v", err)
		c.String(http.StatusBadRequest, err.Error())
		return
	}
	conn, err := larkbase.Connect[Record](c.Request.Context(), os.Getenv("LARK_APP_ID"), os.Getenv("LARK_APP_SECRET"))
	if err != nil {
		log.ErrorLog("fail to connect: %v", err)
		c.String(http.StatusInternalServerError, err.Error())
		return
	}
	id, _ := strconv.Atoi(data.Id)
	var record Record
	err = conn.Find(&record, larkbase.NewFindOption(conn.FilterAnd(conn.Condition().Id.Is(id))))
	if err != nil {
		log.ErrorLog("fail to find record: %v", err)
		c.String(http.StatusInternalServerError, err.Error())
		return
	}
	log.InfoLog("record: %v", conn.MarshalIgnoreError(&record))
	record.Note.SetValue(data.Note)
	err = conn.Update(&record)
	if err != nil {
		log.ErrorLog("fail to update record: %v", err)
		c.String(http.StatusInternalServerError, err.Error())
		return
	}
	log.InfoLog("record updated: %s", record.Id.StringValue())
	c.String(http.StatusOK, "updated")
}

func handleData(c *gin.Context) {
	conn, err := larkbase.Connect[Record](c.Request.Context(), os.Getenv("LARK_APP_ID"), os.Getenv("LARK_APP_SECRET"))
	if err != nil {
		log.ErrorLog("fail to connect: %v", err)
		c.String(http.StatusInternalServerError, err.Error())
		return
	}
	var records []*Record
	err = conn.FindAll(&records, larkbase.NewFindOption(conn.FilterAnd(conn.Condition().Version.Is("v1"))))
	if err != nil {
		log.ErrorLog("fail to find records: %v", err)
		c.String(http.StatusInternalServerError, err.Error())
		return
	}
	fields := []string{"id", "project", "vendor", "sop", "note", "version"}
	data := make([][]string, 0)
	for _, record := range records {
		data = append(data, []string{
			record.Id.StringValue(),
			record.Project.StringValue(),
			record.Vendor.StringValue(),
			record.Sop.StringValue(),
			record.Note.StringValue(),
			record.Version.StringValue(),
		})
	}
	c.JSON(http.StatusOK, gin.H{
		"data":   data,
		"fields": fields,
	})
}
