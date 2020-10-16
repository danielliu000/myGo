package main

import (
	"fmt"
	"strconv"
)

//累加器
//说明:
//testing. appUpper 是一个函数 返回数据类型是 func(int) int
//2. 返回的匿名函数引用了函数体外的n,
//   即返回函数和n形成了引用关系，变成一个整体 构成闭包
//3. 可以理解为: 闭包是类，函数是操作，n是字段
//4. 当反复调用f时，因为n是初始化一次，因此每调用一次就进行累计
//5. 闭包的关键，就是分析出返回的函数使用（引用）了哪些变量
func appUpper() func(int) int {
	var n = 10
	var str = "hello"
	return func (x int) int {
		n += x
		str += strconv.Itoa(x)
		fmt.Println("str=",str)
		return n
	}
}
func main() {
	f := appUpper()
	fmt.Println(f(1))
	fmt.Println(f(2))
	fmt.Println(f(3))
}
