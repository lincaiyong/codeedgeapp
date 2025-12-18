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
	"net/url"
	"os"
	"strconv"
)

func main() {
	codeedgeapp.Run(func(r *gin.RouterGroup) {
		r.POST("/chat/", func(c *gin.Context) {
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
		r.GET("/files/", func(c *gin.Context) {
			params := url.Values{}
			params.Add("project", c.Query("project"))
			params.Add("vendor", c.Query("vendor"))
			b, err := doRequest(fmt.Sprintf("https://codeedge.cc/testeval/files/?%s", params.Encode()))
			if err != nil {
				log.ErrorLog("fail to do request: %v", err)
				c.String(http.StatusInternalServerError, err.Error())
				return
			}
			c.String(http.StatusOK, string(b))
		})
		r.GET("/file/*filepath", func(c *gin.Context) {
			params := url.Values{}
			params.Add("project", c.Query("project"))
			params.Add("vendor", c.Query("vendor"))
			b, err := doRequest(fmt.Sprintf("https://codeedge.cc/testeval/file/%s?%s", c.Param("filepath"), params.Encode()))
			if err != nil {
				log.ErrorLog("fail to do request: %v", err)
				c.String(http.StatusInternalServerError, err.Error())
				return
			}
			c.String(http.StatusOK, string(b))
		})
		r.GET("/search/", func(c *gin.Context) {
			params := url.Values{}
			params.Add("text", c.Query("text"))
			params.Add("flag", c.Query("flag"))
			params.Add("project", c.Query("project"))
			params.Add("vendor", c.Query("vendor"))
			b, err := doRequest(fmt.Sprintf("https://codeedge.cc/testeval/search/?%s", params.Encode()))
			if err != nil {
				log.ErrorLog("fail to do request: %v", err)
				c.String(http.StatusInternalServerError, err.Error())
				return
			}
			c.String(http.StatusOK, string(b))
		})
		type Record struct {
			larkbase.Meta `lark:"https://bytedance.larkoffice.com/base/P8QubLDkzabEJNsaNbacfha0nCd?table=tblUgfvHyAmuS3zx"`
			Id            larkbase.NumberField       `lark:"id"`
			Sop           larkbase.TextField         `lark:"sop"`
			Note          larkbase.TextField         `lark:"note"`
			VulnCode      larkbase.TextField         `lark:"vuln_code"`
			VendorCode    larkbase.TextField         `lark:"vendor_code"`
			Version       larkbase.SingleSelectField `lark:"version"`
		}
		r.POST("/note/", func(c *gin.Context) {
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
		})
		r.GET("/data/", func(c *gin.Context) {
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
			fields := []string{"id", "project", "vendor", "sop", "note"}
			data := make([][]string, 0)
			for _, record := range records {
				data = append(data, []string{
					record.Id.StringValue(),
					record.VulnCode.StringValue(),
					record.VendorCode.StringValue(),
					record.Sop.StringValue(),
					record.Note.StringValue(),
				})
			}
			c.JSON(http.StatusOK, gin.H{
				"data":   data,
				"fields": fields,
			})
		})
	})
}

func doRequest(url string) ([]byte, error) {
	log.InfoLog("doRequest url: %s", url)
	resp, err := http.DefaultClient.Get(url)
	if err != nil {
		return nil, err
	}
	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("http status code: %d", resp.StatusCode)
	}
	defer func() { _ = resp.Body.Close() }()
	b, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	return b, nil
}
