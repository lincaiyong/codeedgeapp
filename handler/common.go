package handler

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

type Response struct {
	Error string `json:"error,omitempty"`
	Data  any    `json:"data,omitempty"`
}

func errorResponse(c *gin.Context, msg string, args ...any) {
	var resp Response
	if len(args) > 0 {
		msg = fmt.Sprintf(msg, args...)
	}
	resp.Error = msg
	c.JSON(http.StatusOK, resp)
}

func dataResponse(c *gin.Context, data any) {
	var resp Response
	resp.Data = data
	c.JSON(http.StatusOK, resp)
}
