package main

import "fmt"

//1. 100以内的数求和，求出当和第一次大于20的当前数
func getNum(num int) int {
	sum := 0
	flag := 0
	for i := 1; i <= num; i++ {
		sum += i
		if sum > 20 {
			flag = i
			break
		}
	}
	return flag
}

//2. 实现登录验证，有三次机会，如果用户名为“张无忌”，密码“888”提示登录成功，否则提示还有几次机会
func login() {
	username := ""
	password := ""
	for i := 1; i <= 3; i++ {
		fmt.Println("输入用户名: ")
		_, _ = fmt.Scanln(&username)
		fmt.Println("输入密码: ")
		_, _ = fmt.Scanln(&password)
		if username == "张无忌" && password == "888" {
			fmt.Println("登录成功")
			break
		} else {
			fmt.Printf("还有%d次机会\n", 3-i)
		}

	}
}

func main() {
	//fmt.Println(getNum(100))
	login()
}
