package handler

import (
	"context"
	"fmt"
	"github.com/lincaiyong/daemon/common"
	"github.com/lincaiyong/log"
	"io"
	"net/http"
	"os"
	"path/filepath"
)

func Download(sample string) {
	url := fmt.Sprintf("https://codeload.github.com/lincaiyong/samples/zip/refs/heads/%s", sample)
	resp, err := http.DefaultClient.Get(url)
	if err != nil {
		log.ErrorLog("fail to download: %v", err)
		return
	}
	defer func() { _ = resp.Body.Close() }()
	if resp.StatusCode != http.StatusOK {
		log.ErrorLog("fail to download: %s", resp.Status)
		return
	}
	b, err := io.ReadAll(resp.Body)
	if err != nil {
		log.ErrorLog("fail to download: %v", err)
		return
	}
	targetPath := filepath.Join("zip", sample+".zip")
	targetPath, _ = filepath.Abs(targetPath)
	err = os.MkdirAll(filepath.Dir(targetPath), os.ModePerm)
	if err != nil {
		log.ErrorLog("fail to create directory: %v", err)
		return
	}
	tmpDir, err := os.MkdirTemp("/tmp", "sample-*")
	if err != nil {
		log.ErrorLog("fail to create temporary directory: %v", err)
		return
	}
	defer func() { _ = os.RemoveAll(tmpDir) }()
	tmpZipFile := filepath.Join(tmpDir, "test.zip")
	err = os.WriteFile(tmpZipFile, b, 0644)
	if err != nil {
		log.ErrorLog("fail to write file: %v", err)
		return
	}
	_, _, err = common.RunCommand(context.Background(), tmpDir, "bash", "-c", fmt.Sprintf("yes | unzip %s", tmpZipFile))
	if err != nil {
		log.ErrorLog("fail to unzip file: %v", err)
		return
	}
	err = os.Remove(tmpZipFile)
	if err != nil {
		log.ErrorLog("fail to remove file: %v", err)
		return
	}
	items, err := os.ReadDir(tmpDir)
	if err != nil {
		log.ErrorLog("fail to unzip file: %v", err)
		return
	}
	if len(items) != 1 || !items[0].IsDir() {
		log.ErrorLog("fail to unzip file")
		return
	}
	tmpDir = filepath.Join(tmpDir, items[0].Name())
	log.InfoLog("source dir: %s", tmpDir)
	_, _, err = common.RunCommand(context.Background(), tmpDir, "bash", "-c", fmt.Sprintf("zip %s -r .", targetPath))
	if err != nil {
		log.ErrorLog("fail to zip file: %v", err)
		return
	}
	log.InfoLog("done")
}
