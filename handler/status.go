package handler

import (
	"github.com/gin-gonic/gin"
	"os"
)

func Status(c *gin.Context) {
	dataResponse(c, map[string]int{
		"pid": os.Getpid(),
	})
}
