package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/lincaiyong/codeedgeapp"
	"github.com/lincaiyong/larkbase"
	"github.com/lincaiyong/log"
	"github.com/lincaiyong/uniapi/service/monica"
	"io"
	"net/http"
	"os"
	"path/filepath"
)

func main() {
	dir, _ := filepath.Abs("..")
	codeedgeapp.Run(func(r *gin.RouterGroup) {
		r.POST("/chat", func(c *gin.Context) {
			var req struct {
				Data string `json:"data"`
			}
			err := c.BindJSON(&req)
			if err != nil {
				log.ErrorLog("fail to bind json: %v", err)
				c.String(http.StatusBadRequest, err.Error())
				return
			}
			c.Header("Access-Control-Allow-Origin", "*")
			c.Header("Access-Control-Expose-Headers", "Content-Type")

			c.Header("Content-Type", "text/event-stream")
			c.Header("Cache-Control", "no-cache")
			c.Header("Connection", "keep-alive")

			monica.Init(os.Getenv("MONICA_SESSION_ID"))
			_, err = monica.ChatCompletion(c.Request.Context(), monica.ModelGPT41Mini, req.Data, func(s string) {
				fmt.Print(s)
				_, _ = fmt.Fprintf(c.Writer, "data: %s\n\n", s)
				c.Writer.Flush()
			})
			if err != nil {
				log.ErrorLog("fail to chat: %v", err)
			}
			_, _ = fmt.Fprintf(c.Writer, "event: close\ndata: \n\n")
			c.Writer.Flush()
		})
		r.GET("/files", func(c *gin.Context) {
			url := fmt.Sprintf("https://codeedge.cc/testeval/files?project=%s&vendor=%s", c.Query("project"), c.Query("vendor"))
			resp, err := http.DefaultClient.Get(url)
			if err != nil {
				log.ErrorLog("fail to request: %v", err)
				c.String(http.StatusInternalServerError, err.Error())
				return
			}
			if resp.StatusCode != http.StatusOK {
				log.ErrorLog("fail to request: code=%d", resp.StatusCode)
				c.String(http.StatusInternalServerError, "fail to request")
				return
			}
			defer func() { _ = resp.Body.Close() }()
			b, err := io.ReadAll(resp.Body)
			if err != nil {
				log.ErrorLog("fail to read body: %v", err)
				c.String(http.StatusInternalServerError, "fail to read body")
				return
			}
			c.String(http.StatusOK, string(b))
		})
		r.GET("/file/*filepath", func(c *gin.Context) {
			filePath := c.Param("filepath")
			b, err := os.ReadFile(filepath.Join(dir, filePath))
			if err != nil {
				log.ErrorLog("fail to read file: %v", err)
				c.String(http.StatusNotFound, err.Error())
				return
			}
			c.String(http.StatusOK, string(b))
		})
		r.GET("/search", func(c *gin.Context) {
			// var args []string
			text := c.Query("text")
			flag := c.Query("flag")
			_, _ = text, flag
			// common.RunCommand(c.Request.Context(), dir, "rg", args...)
			// TODO
			c.Status(http.StatusOK)
		})
		r.GET("/data", func(c *gin.Context) {
			type Record struct {
				larkbase.Meta `lark:"https://bytedance.larkoffice.com/base/P8QubLDkzabEJNsaNbacfha0nCd?table=tblUgfvHyAmuS3zx"`
				Id            larkbase.NumberField `lark:"id"`
				Sop           larkbase.TextField   `lark:"sop"`
				VulnCode      larkbase.TextField   `lark:"vuln_code"`
				SafeCode      larkbase.TextField   `lark:"safe_code"`
			}
			conn, err := larkbase.Connect[Record](c.Request.Context(), os.Getenv("LARK_APP_ID"), os.Getenv("LARK_APP_SECRET"))
			if err != nil {
				log.ErrorLog("fail to connect: %v", err)
				c.String(http.StatusInternalServerError, err.Error())
				return
			}
			var records []*Record
			err = conn.FindAll(&records, larkbase.NewFindOption(conn.FilterAnd(conn.Condition().Id.IsLess(50))))
			if err != nil {
				log.ErrorLog("fail to find records: %v", err)
				c.String(http.StatusInternalServerError, err.Error())
				return
			}
			fields := []string{"id", "project", "note"}
			data := make([][]string, 0)
			for _, record := range records {
				data = append(data, []string{
					record.Id.StringValue(),
					record.VulnCode.StringValue(),
					record.Sop.StringValue(),
				})
			}
			c.JSON(http.StatusOK, gin.H{
				"data":   data,
				"fields": fields,
			})
		})
	})
}
