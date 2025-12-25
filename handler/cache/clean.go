package cache

import (
	"context"
	"github.com/lincaiyong/daemon/common"
	"github.com/lincaiyong/log"
)

var deleteFlag int

func deleteEmptyDirectory(ctx context.Context) {
	if deleteFlag > 0 {
		deleteFlag--
		return
	}
	deleteFlag = 100
	_, _, err := common.RunCommand(ctx, cacheDir, "find", ".", "-type", "d", "-empty", "-depth", "-delete")
	if err != nil {
		log.WarnLog("fail to clean: %v", err)
	}
}
