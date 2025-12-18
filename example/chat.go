package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/lincaiyong/log"
	"github.com/lincaiyong/uniapi/service/monica"
	"net/http"
	"os"
)

func handleChat(c *gin.Context) {
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
}
