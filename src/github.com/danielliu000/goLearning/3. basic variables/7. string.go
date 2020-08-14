package main

import "fmt"

func main() {
	// "xxx" 正常双引号声明字符串，但双引号处理特殊字符能力差
	// `xxx` 反引号声明字符串 类似es6

	str := `hello, world`
	fmt.Printf(`str的类型是: %T` + "\n", str)
	//支持切片
	fmt.Println(str[:])
	fmt.Println(str[:5])
	fmt.Println(str[7:])
}
