package lru_file

import (
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


    buf := make([]byte, 102400)
    for {
        n, _ := file.Read(buf)
        if n == 0 {
            break
        }
        _, error := out.Write(buf[:n])
        if error != nil {
            return "", "", nil, error
        }
    }
    // 这个方法适合写入小文件, 因为方法内部会强制开辟一块和文件一样大小的内存
    //_, err = io.Copy(out, file)

    fileInfo, err := os.Stat(temporaryPath + string(os.PathSeparator) + fileName)
    if err != nil {
       return "", "", nil, err
    }

    return fileName, temporaryPath + string(os.PathSeparator) + fileName, fileInfo, nil
}




