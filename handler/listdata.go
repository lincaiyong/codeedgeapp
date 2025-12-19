package handler

import (
	"github.com/gin-gonic/gin"
	"sort"
)

func ListData(c *gin.Context) {
	var result []string
	for k := range conf.DataUrl {
		result = append(result, k)
	}
	sort.Strings(result)
	dataResponse(c, result)
}
