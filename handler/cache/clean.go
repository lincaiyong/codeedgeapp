package cache

import (
	"context"
	"github.com/lincaiyong/daemon/common"
	"github.com/lincaiyong/log"
)

func deleteEmptyDirectory(ctx context.Context) {
	_, _, err := common.RunCommand(ctx, cacheDir, "find", ".", "-type", "d", "-empty", "-depth", "-delete")
	if err != nil {
		log.WarnLog("fail to clean: %v", err)
	}
}
