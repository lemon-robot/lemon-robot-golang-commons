package lru_string

import (
	"github.com/satori/go.uuid"
	"strings"
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

func (i *LRUString) Uuid(withoutLine bool) string {
	if withoutLine {
		return strings.ReplaceAll(uuid.Must(uuid.NewV4()).String(), "-", "")
	}
	return uuid.Must(uuid.NewV4()).String()
}
