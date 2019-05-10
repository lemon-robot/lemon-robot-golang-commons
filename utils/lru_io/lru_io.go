package lru_io

import (
	"encoding/json"
	"io"
	"io/ioutil"
	"os"
	"path"
	"path/filepath"
	"strings"
	"sync"
)

type LRUIO struct {
}

var instance *LRUIO
var once sync.Once

func GetInstance() *LRUIO {
	once.Do(func() {
		instance = &LRUIO{}
	})
	return instance
}

func (i *LRUIO) PathExists(path string) bool {
	_, err := os.Stat(path)
	if err == nil {
		return true
	}
	if os.IsNotExist(err) {
		return false
	}
	return false
}

func (i *LRUIO) CopyFile(src string, dst string) error {
	srcFile, errSrc := os.Open(src)
	if errSrc != nil {
		return errSrc
	}
	return i.CopyFileFromReader(srcFile, dst)
}

func (i *LRUIO) CopyFileFromReader(srcFile io.Reader, dst string) error {
	dirPath, _ := path.Split(dst)
	if !i.PathExists(dirPath) {
		dstDirErr := os.MkdirAll(dirPath, os.ModePerm)
		if dstDirErr != nil {
			return dstDirErr
		}
	}
	dstFile, errDstCreate := os.Create(dst)
	if errDstCreate != nil {
		return errDstCreate
	}
	defer dstFile.Close()
	_, errDestCopy := io.Copy(dstFile, srcFile)
	if errDestCopy != nil {
		return errDestCopy
	}
	return nil
}

func (i *LRUIO) CopyDir(src string, dest string) error {
	err := filepath.Walk(src, func(currentSrc string, f os.FileInfo, err error) error {
		if f == nil {
			return err
		}
		aimPath := strings.Replace(currentSrc, src, dest, 1)
		if !f.IsDir() {
			i.CopyFile(currentSrc, aimPath)
		}
		return nil
	})
	if err != nil {
		return err
	}
	return nil
}

func (i *LRUIO) JsonToStruct(jsonSrc string, obj interface{}) error {
	data, readErr := ioutil.ReadFile(jsonSrc)
	if readErr != nil {
		return readErr
	}
	parseErr := json.Unmarshal(data, obj)
	if parseErr != nil {
		return parseErr
	}
	return nil
}

func (i *LRUIO) GetRuntimePath(filename string) string {
	dir, _ := filepath.Abs(filepath.Dir(os.Args[0]))
	return filepath.Join(dir, filename)
}

func (i *LRUIO) ReplaceStrToFile(content, path string) error {
	if i.PathExists(path) {
		err := os.Remove(path)
		if err != nil {
			return err
		}
	}
	file, err := os.Create(path)
	if err != nil {
		return err
	}
	defer file.Close()
	_, wErr := file.WriteString(content)
	if wErr != nil {
		return wErr
	}
	return nil
}
