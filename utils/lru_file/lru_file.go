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

func (i *LRUFile) SaveFileToTemporary(file multipart.File, handler *multipart.FileHeader) (string, string, os.FileInfo, error) {
    fileName := handler.Filename
    out, err := os.Create(temporaryPath + string(os.PathSeparator) + fileName)
    if err != nil {
        return "", "", nil, err
    }
    defer out.Close()
    _, err = io.Copy(out, file)
    if err != nil {
        return "", "", nil, err
    }
    fileInfo, err := os.Stat(temporaryPath + string(os.PathSeparator) + fileName)
    if err != nil {
        return "", "", nil, err
    }
    return fileName, temporaryPath + string(os.PathSeparator) + fileName, fileInfo, nil
}



