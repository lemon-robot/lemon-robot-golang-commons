package lru_machine

import (
	"encoding/json"
	"lemon-robot-golang-commons/logger"
	"lemon-robot-golang-commons/model"
	"lemon-robot-golang-commons/utils/lru_io"
	"lemon-robot-golang-commons/utils/lru_string"
	"lemon-robot-golang-commons/utils/lru_workspace"
	"os"
	"sync"
)

type LRUMachine struct {
	machineSignCache string
}

const machineSignSaveFile = "lr.msign"

var instance *LRUMachine
var once sync.Once

func GetInstance() *LRUMachine {
	once.Do(func() {
		instance = &LRUMachine{}
	})
	return instance
}

func (i *LRUMachine) GetMachineSign() string {
	if i.machineSignCache == "" {
		machineSign, err := i.initMachineSign()
		if err != nil {
			logger.Error("Server nodes cannot be registered because lru_machine sign cannot be computed", err)
			os.Exit(1)
			return ""
		}
		i.machineSignCache = machineSign
	}
	return i.machineSignCache
}

func (i *LRUMachine) initMachineSign() (string, error) {
	machineSign := i.readMachineSignFromLocal()
	if machineSign == "" {
		logger.Info("Machine Sign does not exist locally, start generating...")
		machineSign = lru_string.GetInstance().Uuid()
		saveErr := i.saveMachineSignToLocal(machineSign)
		if saveErr != nil {
			return "", saveErr
		}
		logger.Info("Machine Sign generation is completed: " + machineSign)
	}
	return machineSign, nil
}

func (i *LRUMachine) saveMachineSignToLocal(machineSign string) error {
	signObj := &model.LrMachineSign{
		Sign: machineSign,
	}
	jsonBytes, _ := json.Marshal(signObj)
	return lru_io.GetInstance().ReplaceStrToFile(string(jsonBytes), lru_workspace.GetInstance().GetWorkspacePath(machineSignSaveFile))
}

func (i *LRUMachine) readMachineSignFromLocal() string {
	signObj := model.LrMachineSign{}
	err := lru_io.GetInstance().JsonToStruct(lru_workspace.GetInstance().GetWorkspacePath(machineSignSaveFile), &signObj)
	if err != nil {
		return ""
	}
	return signObj.Sign
}
