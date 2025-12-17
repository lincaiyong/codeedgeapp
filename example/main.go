package main

import (
	"github.com/gin-gonic/gin"
	"github.com/lincaiyong/codeedgeapp"
	"github.com/lincaiyong/log"
	"io/fs"
	"net/http"
	"os"
	"path/filepath"
	"strings"
)

func main() {
	dir, _ := filepath.Abs("..")
	codeedgeapp.Run(func(r *gin.RouterGroup) {
		r.GET("/files", func(c *gin.Context) {
			result := make([]string, 0)
			err := filepath.WalkDir(dir, func(path string, d fs.DirEntry, err error) error {
				if err != nil {
					return err
				}
				if d.IsDir() && strings.HasPrefix(d.Name(), ".") {
					return filepath.SkipDir
				}
				if !d.IsDir() {
					result = append(result, path[len(dir)+1:])
				}
				return nil
			})
			if err != nil {
				log.ErrorLog("fail to walk dir: %v", err)
				c.String(http.StatusInternalServerError, err.Error())
				return
			}
			c.JSON(http.StatusOK, result)
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
	})
}
