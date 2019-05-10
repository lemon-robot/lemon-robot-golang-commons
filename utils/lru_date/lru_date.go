package lru_date

import (
	"sync"
	"time"
)

type LRUDate struct {
}

var instance *LRUDate
var once sync.Once

func GetInstance() *LRUDate {
	once.Do(func() {
		instance = &LRUDate{}
	})
	return instance
}

func GetCurrentTimeFormatStr() string {
	return time.Now().Format("2006-01-02 15:04:05")
}
