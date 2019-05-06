package workspace

import (
	"fmt"
	"lemon-robot-golang-commons/logger"
	"lemon-robot-golang-commons/utils/lruio"
	"os"
)

const workspaceDirName = "lr_workspace"

func GetWorkspacePath(fileName string) string {
	workspaceDirFullPath := lruio.GetRuntimePath(workspaceDirName)
	if !lruio.PathExists(workspaceDirFullPath) {
		err := os.MkdirAll(workspaceDirFullPath, os.ModePerm)
		if err != nil {
			logger.Error("Cannot init workspace dir at : "+workspaceDirFullPath, err)
			os.Exit(1)
		}
	}
	return lruio.GetRuntimePath(fmt.Sprintf("%s/%s", workspaceDirName, fileName))
}
