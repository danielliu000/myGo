package account

import (
	"fmt"
	"strings"
)

func (a *account) Exit() bool {
	keyPress := ""
	for {
		fmt.Printf("你确定要退出吗? y/n: ")
		_, _ = fmt.Scanln(&keyPress)
		if strings.ToLower(keyPress) == "y" {
			fmt.Println()
			fmt.Println("程序退出, 再见")
			return true
		} else if strings.ToLower(keyPress) == "n" {
			return false
		} else {
			fmt.Println("请输入正确的指令: y or n")
		}
	}

}