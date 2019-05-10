package lru_string

import (
	"github.com/satori/go.uuid"
	"sync"
)

type LRUString struct{}

var instance *LRUString
var once sync.Once

func GetInstance() *LRUString {
	once.Do(func() {
		instance = &LRUString{}
	})
	return instance
}

func (i *LRUString) Uuid() string {
	return uuid.Must(uuid.NewV4()).String()
}
