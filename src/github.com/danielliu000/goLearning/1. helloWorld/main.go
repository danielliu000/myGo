package main  // 每个go文件都必须归属一个包

import "fmt"

func main() {
	//fmt.Println("hello world")
	fmt.Println("hello world\rGo")// 回车但不换行 hello world 会被替换成Go


}

//go run xxx.go 编译成可执行文件 默认使用xxx命名