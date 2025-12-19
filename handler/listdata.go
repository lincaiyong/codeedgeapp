package handler

import (
	"github.com/gin-gonic/gin"
	"sort"
)

func ListData(c *gin.Context) {
	var result []string
	for k := range config.DataUrl {
		result = append(result, k)
	}
	sort.Strings(result)
	dataResponse(c, result)
}
