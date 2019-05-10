package lru_cmd

import (
	"os"
	"os/exec"
	"sync"
)

type LRUCmd struct {
}

var instance *LRUCmd
var once sync.Once

func GetInstance() *LRUCmd {
	once.Do(func() {
		instance = &LRUCmd{}
	})
	return instance
}

func (i *LRUCmd) NoDisplayExec(command string) error {
	cmd := exec.Command("/bin/bash", "-c", command)
	_, err := cmd.Output()
	return err
}

func (i *LRUCmd) OnlyConsoleDisplayExec(command string) {
	cmd := exec.Command("/bin/bash", "-c", command)
	cmd.Stdout = os.Stdout //
	cmd.Run()
}
