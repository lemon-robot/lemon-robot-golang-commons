package lru_string

import (
    "errors"
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


func (i *LRUString)GetFileSuffixName(fileName string) (error, string) {
    var fileFormatArr = [...]string{".cpio.gz", ".cpio.cgz", ".tar.gz", ".tar.bz", ".tar.bz2", ".tar.xz", ".tar.z"}
    for _, v := range fileFormatArr {
        if strings.Index(fileName, v) != -1 {
            return nil, v
        }
    }
    if strings.Index(fileName, ".") == -1 {
        return errors.New("Unable to get file type"), ""
    }
    endIndex := strings.LastIndex(fileName, ".")
    fileType := fileName[endIndex : len(fileName)]
    return nil, fileType
}
