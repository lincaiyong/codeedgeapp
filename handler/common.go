package handler

import (
	"archive/zip"
	"context"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/lincaiyong/daemon/common"
	"io"
	"net/http"
	"os"
	"path/filepath"
	"strings"
	"time"
)

type Response struct {
	Error string `json:"message,omitempty"`
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

func readZipFiles(zipPath string) ([]string, error) {
	reader, err := zip.OpenReader(zipPath)
	if err != nil {
		return nil, fmt.Errorf("打开zip文件失败: %w", err)
	}
	defer func() { _ = reader.Close() }()
	var files []string
	for _, file := range reader.File {
		if file.FileInfo().IsDir() {
			continue
		}
		files = append(files, file.Name)
	}
	return files, nil
}

func readZipFile(zipPath, fileName string) ([]byte, error) {
	reader, err := zip.OpenReader(zipPath)
	if err != nil {
		return nil, err
	}
	defer func() { _ = reader.Close() }()
	var file *zip.File
	for _, f := range reader.File {
		name := f.Name
		if name == fileName {
			file = f
			break
		}
	}
	if file == nil {
		return nil, fmt.Errorf("file not found: %s", fileName)
	}
	rc, err := file.Open()
	if err != nil {
		return nil, err
	}
	defer func() { _ = rc.Close() }()
	b, err := io.ReadAll(rc)
	if err != nil {
		return nil, err
	}
	return b, nil
}

func modifiedTime(filePath string) (time.Time, error) {
	stat, err := os.Stat(filePath)
	if err != nil {
		return time.Time{}, err
	}
	return stat.ModTime(), nil
}

func unzipSample(zipPath, sample string) (string, error) {
	dirPath := fmt.Sprintf("/tmp/sample-%s", strings.ReplaceAll(sample, "/", "-"))
	_ = os.RemoveAll(dirPath)
	if stat, err := os.Stat(dirPath); err == nil && stat.IsDir() {
		items, _ := os.ReadDir(dirPath)
		if len(items) > 0 {
			return dirPath, nil
		}
		_ = os.RemoveAll(dirPath)
	}
	err := os.MkdirAll(dirPath, os.ModePerm)
	if err != nil {
		return "", err
	}
	zipPath, _ = filepath.Abs(zipPath)
	_, _, err = common.RunCommand(context.Background(), dirPath, "bash", "-c", fmt.Sprintf("yes | unzip %s || true", zipPath))
	if err != nil {
		return "", err
	}
	return dirPath, nil
}
