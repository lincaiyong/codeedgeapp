package handler

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/lincaiyong/codeedgeapp/handler/cache"
	"github.com/lincaiyong/daemon/common"
	"github.com/lincaiyong/gui"
	"net/http"
	"path/filepath"
	"strings"
)

func Diff(c *gin.Context) {
	lhs := c.Query("lhs")
	rhs := c.Query("rhs")
	if lhs == "" || rhs == "" {
		errorResponse(c, "invalid lhs or rhs")
		return
	}
	mod, err := cache.GetModTime(lhs)
	if err != nil {
		errorResponse(c, "project not found: %v", err)
		return
	}
	if gui.IfNotModifiedSince(c, mod) {
		c.String(http.StatusNotModified, "not modified")
		return
	}
	lhsDir, err := cache.EnsureProjectDir(lhs)
	if err != nil {
		errorResponse(c, "project not found: %v", err)
		return
	}
	rhsDir, err := cache.EnsureProjectDir(rhs)
	if err != nil {
		errorResponse(c, "project not found: %v", err)
		return
	}
	args := []string{
		"-c", fmt.Sprintf("diff -rq %s %s | grep differ | awk '{print $2}'", lhsDir, rhsDir),
	}
	stdout, _, err := common.RunCommand(c.Request.Context(), "", "bash", args...)
	if err != nil {
		errorResponse(c, "fail to run diff: %v", err)
		return
	}
	result := make([]string, 0)
	if stdout != "" {
		for _, line := range strings.Split(strings.TrimSpace(stdout), "\n") {
			relPath, _ := filepath.Rel(lhsDir, line)
			result = append(result, relPath)
		}
	}
	gui.SetLastModified(c, mod, 0)
	dataResponse(c, result)
}
