package handler

import (
	_ "embed"
	"github.com/gin-gonic/gin"
	"net/http"
	"regexp"
)

func GetObject(c *gin.Context) {
	hash := c.Param("sha1")
	if !regexp.MustCompile(`^[a-f0-9]+$`).MatchString(hash) {
		c.String(http.StatusBadRequest, "hash is invalid: "+hash)
		return
	}
	data, err := conf.ObjectFn(c.Request.Context(), hash)
	if err != nil {
		c.String(http.StatusNotFound, "object not found")
		return
	}
	c.Header("Content-Disposition", "attachment; filename="+hash)
	c.Header("Content-Type", "application/octet-stream")
	c.Data(http.StatusOK, "application/octet-stream", data)
}
