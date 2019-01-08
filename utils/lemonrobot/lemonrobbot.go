package lemonrobot

import "fmt"

func PrintInfo(appName, appVersion string) {
	fmt.Println(`
 _                               ______       _                
| |                             (_____ \     | |           _   
| |      _____ ____   ___  ____  _____) )___ | |__   ___ _| |_ 
| |     | ___ |    \ / _ \|  _ \|  __  // _ \|  _ \ / _ (_   _)
| |_____| ____| | | | |_| | | | | |  \ \ |_| | |_) ) |_| || |_ 
|_______)_____)_|_|_|\___/|_| |_|_|   |_\___/|____/ \___/  \__)

	`)
	fmt.Printf("-- %c[0;0;32m%s - Version: %s - https://www.lemonit.cn%c[0m --\n\n", 0x1B, appName, appVersion, 0x1B)
}
