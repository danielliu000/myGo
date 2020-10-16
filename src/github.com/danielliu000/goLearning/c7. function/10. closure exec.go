package main

//testing. 编写一个函数makeSuffix(suffix string)
//可以接收一个文件后缀名(比如.jpg)，并返回一个闭包
//2. 调用闭包，可以传入一个文件名，
//如果该文件名没有指定的后缀，则返回文件名.jpg
//如果已经有.jpg后缀，则返回原文件名
//3. 提示: strings.HasSuffix
import (
	"fmt"
	"strings"
)

func makeSuffix(suffix string) func(string) string {

	return func(filename string) string {
		if !strings.HasSuffix(filename, suffix) {
			return filename + suffix
		}
		return filename
	}
}

func main() {
	f := makeSuffix(".jpg")

	fmt.Println(f("xxx.jpg"))
	fmt.Println(f("yyy"))
}
