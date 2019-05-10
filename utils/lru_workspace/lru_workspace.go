package lru_workspace

import (
	"fmt"
	"lemon-robot-golang-commons/logger"
	"lemon-robot-golang-commons/utils/lru_io"
	"os"
	"sync"
)

type LRUWorkspace struct{}

const workspaceDirName = "lr_workspace"

var instance *LRUWorkspace
var once sync.Once

func (i *LRUWorkspace) GetInstance() *LRUWorkspace {
	once.Do(func() {
		instance = &LRUWorkspace{}
	})
	return instance
}

func (i *LRUWorkspace) GetWorkspacePath(fileName string) string {
	workspaceDirFullPath := lru_io.GetInstance().GetRuntimePath(workspaceDirName)
	if !lru_io.GetInstance().PathExists(workspaceDirFullPath) {
		err := os.MkdirAll(workspaceDirFullPath, os.ModePerm)
		if err != nil {
			logger.Error("Cannot init lru_workspace dir at : "+workspaceDirFullPath, err)
			os.Exit(1)
		}
	}
	return lru_io.GetInstance().GetRuntimePath(fmt.Sprintf("%s/%s", workspaceDirName, fileName))
}
