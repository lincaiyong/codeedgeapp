package handler

import (
	"codeedgeapp/handler/cache"
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/lincaiyong/daemon/common"
	"github.com/lincaiyong/gui"
	"github.com/lincaiyong/log"
	"net/http"
	"sort"
	"strings"
	"time"
)

func Search(c *gin.Context) {
	project := c.Query("project")
	vendor := c.Query("vendor")
	if project == "" || strings.Contains(project, ".") || strings.Contains(vendor, ".") {
		errorResponse(c, "project or vendor is invalid")
		return
	}
	text := c.Query("text")
	flag := c.Query("flag")
	if strings.TrimSpace(text) == "" {
		errorResponse(c, "text is invalid")
		return
	}
	projectMod, err := cache.GetModTime(project)
	if err != nil {
		errorResponse(c, "fail to get modified time: %v", err)
		return
	}
	if vendor == "" && gui.IfNotModifiedSince(c, projectMod) {
		c.String(http.StatusNotModified, "not modified")
		return
	}
	projects := []string{project}
	if vendor != "" {
		for _, item := range strings.Split(vendor, ",") {
			var vendorMod time.Time
			vendorMod, err = cache.GetModTime(item)
			if err != nil {
				errorResponse(c, "fail to get modified time: %v", err)
				return
			}
			if projectMod.Before(vendorMod) {
				projectMod = vendorMod
			}
			projects = append(projects, item)
		}
	}
	if gui.IfNotModifiedSince(c, projectMod) {
		c.String(http.StatusNotModified, "not modified")
		return
	}
	var result []*RipgrepItem
	for _, item := range projects {
		var tmp []*RipgrepItem
		tmp, err = searchProject(c, item, text, flag)
		if err != nil {
			errorResponse(c, "fail to search project: %v", err)
			return
		}
		if item != project {
			for _, t := range tmp {
				t.Path = fmt.Sprintf("@vendor/%s/%s", strings.ReplaceAll(item, "/", "-"), t.Path)
			}
		}
		result = append(result, tmp...)
	}
	b, _ := json.MarshalIndent(result, "", "    ")
	gui.SetLastModified(c, projectMod, 0)
	c.String(http.StatusOK, string(b))
}

func searchProject(c *gin.Context, project, text, flag string) ([]*RipgrepItem, error) {
	workDir, err := cache.EnsureProjectDir(project)
	if err != nil {
		return nil, err
	}
	log.InfoLog("workDir: %s", workDir)
	args := strings.Fields(flag)
	args = append(args, "--json", text)
	stdout, _, err := common.RunCommand(c.Request.Context(), workDir, "rg", args...)
	var result []*RipgrepItem
	resultMap := make(map[string]int)
	var keys []string
	if err == nil {
		var unsorted []*RipgrepItem
		for _, line := range strings.Split(stdout, "\n") {
			var rgLine RipgrepLine
			if err = json.Unmarshal([]byte(line), &rgLine); err != nil {
				log.ErrorLog("fail to unmarshal rg: %v", err)
				continue
			}
			if rgLine.Type == "match" {
				var matchIndex []int
				for _, m := range rgLine.Data.Submatches {
					matchIndex = append(matchIndex, m.Start, m.End)
				}
				item := &RipgrepItem{
					Path:       rgLine.Data.Path.Text,
					LineText:   strings.TrimRight(rgLine.Data.Lines.Text, "\n"),
					LineNumber: rgLine.Data.LineNumber,
					MatchIndex: matchIndex,
				}
				key := fmt.Sprintf("%s:%d", item.Path, item.LineNumber)
				keys = append(keys, key)
				resultMap[key] = len(unsorted)
				unsorted = append(unsorted, item)
				if len(unsorted) > 100 {
					break
				}
			}
		}
		sort.Strings(keys)
		for _, key := range keys {
			result = append(result, unsorted[resultMap[key]])
		}
	} else {
		result = []*RipgrepItem{}
		log.ErrorLog("fail to run rg: %v", err)
	}
	return result, nil
}

type RipgrepLine struct {
	Type string `json:"type"`
	Data struct {
		Path struct {
			Text string `json:"text"`
		} `json:"path"`
		Lines struct {
			Text string `json:"text"`
		} `json:"lines"`
		LineNumber int `json:"line_number"`
		Submatches []struct {
			Match struct {
				Text string `json:"text"`
			} `json:"match"`
			Start int `json:"start"`
			End   int `json:"end"`
		} `json:"submatches"`
	} `json:"data"`
}

type RipgrepItem struct {
	Path       string `json:"path"`
	LineText   string `json:"line_text"`
	LineNumber int    `json:"line_number"`
	MatchIndex []int  `json:"match_index"`
}
