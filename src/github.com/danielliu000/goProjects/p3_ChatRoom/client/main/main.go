package main

import (
	"chatRoom/client/process"
	"fmt"
	"os"
)

var userId int
var userPwd string
var userName string

func main() {
	var key int
	//var loop = true
	for {
		fmt.Println("-------------欢迎登录多人聊天系统-------------")
		fmt.Println("\t\t\t 1 登录聊天系统")
		fmt.Println("\t\t\t 2 注册用户")
		fmt.Println("\t\t\t 3 退出系统")

		fmt.Println("请选择（1-3）：")
		fmt.Println("---------------------------------------------")

		_, _ = fmt.Scanf("%d\n", &key)
		switch key {
		case 1:
			fmt.Println("登录")
			fmt.Println("请输入用户ID：")
			_, _ = fmt.Scanln(&userId)
			fmt.Println("请输入用户密码： ")
			_, _ = fmt.Scanln(&userPwd)
			//loop = false
			//完成登录
			//1.创建userProcess实例

			up := &process.UserProcess{}
			_ = up.Login(userId, userPwd)

		case 2:
			fmt.Println("注册用户")
			fmt.Println("请输入用户ID：")
			_, _ = fmt.Scanf("%d\n", &userId)
			fmt.Println("请输入用户密码：")
			_, _ = fmt.Scanf("%s\n", &userPwd)
			fmt.Println("请输入用户名：")
			_, _ = fmt.Scanf("%s\n", &userName)
			up := &process.UserProcess{}
			_ = up.Register(userId, userPwd, userName)
			//loop = false
		case 3:
			fmt.Println("退出系统")
			os.Exit(0)

		default:
			fmt.Println("你的输入有误，请重新输入")
		}
	}



}