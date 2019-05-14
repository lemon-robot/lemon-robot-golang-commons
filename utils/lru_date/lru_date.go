package lru_date

import (
	"fmt"
	"lemon-robot-server/sysinfo"
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

func (i *LRUDate) GetCurrentTimeFormatStr() string {
	return time.Now().Format("2006-01-02 15:04:05")
}

func (i *LRUDate) CalculateTimeByDurationStr(preTime time.Time, durStr string) time.Time {
	dur, _ := time.ParseDuration(fmt.Sprintf("-%ds", sysinfo.LrServerConfig().ClusterNodeActiveInterval*2))
	return preTime.Add(dur)
}
