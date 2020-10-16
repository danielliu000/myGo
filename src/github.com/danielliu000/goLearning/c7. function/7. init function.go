package main

import (
	"fmt"
	"utils" //init()最先执行 // testing
)

//每个go文件都可以定义一个init函数，该函数在main函数调用前执行,类似构造函数
//testing. 如果一个文件同时包含全局变量定义，init函数，main函数，则执行流程是
//	变量定义 => init函数 => main函数
//2. 如果main.go 和 utils.go都有变量定义,init函数，执行顺序是:
//  utils.go ->变量定义 -> init函数 => main.go - 变量定义 -> init函数 -> main函数

var myAge = getAge()

func getAge() int {//2
	fmt.Println("getAge()")
	return 90
}
func init() {//3
	fmt.Println("i am an init c7. function")
}
func main() { //4
	fmt.Println("i am a main c7. function, age=",myAge)
	fmt.Println("age=", utils.Age)
	fmt.Println("name=", utils.Name)
}
