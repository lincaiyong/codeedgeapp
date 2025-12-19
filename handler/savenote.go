package handler

import (
	"github.com/gin-gonic/gin"
)

func SaveNote(c *gin.Context) {
	//var data struct {
	//	Note string `json:"note"`
	//	Id   string `json:"id"`
	//}
	//err := c.BindJSON(&data)
	//if err != nil {
	//	log.ErrorLog("fail to bind json: %v", err)
	//	c.String(http.StatusBadRequest, err.Error())
	//	return
	//}
	//conn, err := larkbase.ConnectAny(c.Request.Context(), os.Getenv("LARK_APP_ID"), os.Getenv("LARK_APP_SECRET"))
	//if err != nil {
	//	log.ErrorLog("fail to connect: %v", err)
	//	c.String(http.StatusInternalServerError, err.Error())
	//	return
	//}
	//id, _ := strconv.Atoi(data.Id)
	//var record Record
	//err = conn.Find(&record, larkbase.NewFindOption(conn.FilterAnd(conn.Condition().Id.Is(id))))
	//if err != nil {
	//	log.ErrorLog("fail to find record: %v", err)
	//	c.String(http.StatusInternalServerError, err.Error())
	//	return
	//}
	//log.InfoLog("record: %v", conn.MarshalIgnoreError(&record))
	//record.Note.SetValue(data.Note)
	//err = conn.Update(&record)
	//if err != nil {
	//	log.ErrorLog("fail to update record: %v", err)
	//	c.String(http.StatusInternalServerError, err.Error())
	//	return
	//}
	//log.InfoLog("record updated: %s", record.Id.StringValue())
	//c.String(http.StatusOK, "updated")
	errorResponse(c, "not implemented")
}
