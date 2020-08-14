package main

import (
	"fmt"
	"os"
)
//当执行到defer时，暂不执行，会将defer后的语句压入到独立的defer栈中
//当函数执行完毕，再从defer栈中，按照先入后出的方式出栈 和执行
//如果入栈时涉及到变量，则将变量的值同时入栈（保存）出栈时取出保存的值
// 即时用指针 值也不会变
func sum(n1 *int, n2 *int) int {
	defer fmt.Println("ok n1=",*n1) // 3
	defer fmt.Println("ok n2=",*n2) //2
	*n1++ //不会影响上面defer中的n1
	*n2++ //不会影响上面defer中的n2
	res := *n1 + *n2
	fmt.Println("ok res=", res)  //1
	return res
}

//最佳实践 关闭文件或关闭资源时使用

func optFiles() {
	file = openfile("filename")
	defer file.close() // 函数执行完 关闭文件 不知道啥时候用完，反正 用完了defer会处理
}
func optData() {
	connect = openDB()
	defer connect.close() //非常好用，不知道啥时候关闭数据库， 反正会defer会关
}
func main() {
	n1 := 20
	n2 := 30
	res := sum(&n1, &n2)
	fmt.Println("res", res) //4
}
