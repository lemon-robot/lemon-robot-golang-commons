package lrumachine

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"net"
	"strings"
)

// 通过Mac地址计算机器码
func CalculateMachineCodeByMAC() (string, error) {
	interfaces, err := net.Interfaces()
	if err != nil {
		return "", err
	}
	var macAddrs []string = make([]string, len(interfaces))
	i := 0
	for _, inter := range interfaces {
		macAddrs[i] = fmt.Sprintf("%v", inter.HardwareAddr)
	}
	md5Obj := md5.New()
	md5Obj.Write([]byte(strings.Join(macAddrs, "-")))
	return hex.EncodeToString(md5Obj.Sum(nil)), nil
}
