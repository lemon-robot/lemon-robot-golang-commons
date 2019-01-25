package lrucmd

import (
	"os"
	"os/exec"
)

func NoDisplayExec(command string) error {
	cmd := exec.Command("/bin/bash", "-c", command)
	_, err := cmd.Output()
	return err
}

func OnlyConsoleDisplayExec(command string) {
	cmd := exec.Command("/bin/bash", "-c", command)
	cmd.Stdout = os.Stdout //
	cmd.Run()
}
