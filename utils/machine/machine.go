package lrumachine

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"net"
	"runtime"
	"strings"
)

// 通过Mac地址计算机器码
func CalculateMachineCode() (string, error) {
	interfaces, err := net.Interfaces()
	if err != nil {
		return "", err
	}
	var items = make([]string, len(interfaces)+2)
	i := 0
	for _, inter := range interfaces {
		items[i] = fmt.Sprintf("%v", inter.HardwareAddr)
		i += 1
	}
	items[i] = runtime.GOOS
	items[i+1] = runtime.GOARCH
	md5Obj := md5.New()
	md5Obj.Write([]byte(strings.Join(items, "")))
	return hex.EncodeToString(md5Obj.Sum(nil)), nil
}
