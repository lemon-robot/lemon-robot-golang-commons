package lru_file

import (
    "io"
    "mime/multipart"
    "os"
    "sync"
)

type LRUFile struct {

}

var instance *LRUFile
var once sync.Once

var temporaryPath = os.TempDir()

func GetInstance() *LRUFile {
    once.Do(func() {
        instance = &LRUFile{}
    })
    return instance
}

func (i *LRUFile) SaveFileToTemporary(file multipart.File, handler *multipart.FileHeader) (error, string, string) {
    fileName := handler.Filename
    out, err := os.Create(temporaryPath + fileName)
    if err != nil {
        return err, "", ""
    }
    defer out.Close()
    _, err = io.Copy(out, file)
    if err != nil {
        return err, "", ""
    }
    return nil, fileName, temporaryPath + fileName
}



