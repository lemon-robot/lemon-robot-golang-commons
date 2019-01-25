package logger

import (
	"fmt"
	"lemon-robot-golang-commons/utils/lrudate"
)

func Debug(msg string) {
	log(msg, 0)
}

func Warn(msg string) {
	log(msg, 1)
}

func Error(msg string, err error) {
	log(msg, 2)
	if err != nil {
		log(err.Error(), 2)
	}
}

func Info(msg string) {
	log(msg, 3)
}

var logTypeList = []string{"DEBG", "WARN", "ERRO", "INFO"}
var logColorList = []int{0, 33, 31, 36}

func log(msg string, logType int) {
	fmt.Printf("%c[1;0;%dm[%s %s]%c[0m %s\n", 0x1B, logColorList[logType], logTypeList[logType], lrudate.GetCurrentTimeFormatedStr(), 0x1B, msg)
}
