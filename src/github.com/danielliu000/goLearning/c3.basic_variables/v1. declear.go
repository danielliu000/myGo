package main

import (
	"fmt"
)

//var m = 30  //全局中 不可以使用:=

func main() {
	//testing . 声明变量  var 变量名 变量类型
	//var name string // default ''
	//var age uint8 // default 0
	//var price float32 // default 0
	//var isOk bool // default false

	//2. 每个变量写一个var太繁琐 批量声明
	//var(
	//	name string = "Daniel"
	//	age int = 34
	//	isOk bool = true
	//)

	//3. 省略类型方式声明变量，此时编译器会匹配赋值类型
	//var (
	//	name = "Winnie"
	//	age = 33
	//	isOk = false
	//)

	//4. 语法糖 省略var 和 类型 来声明变量 := 仅在函数内部使用
	//name := "Leo"
	//age := 33
	//price:=3.33
	//isOk := true
	//m := 50  // 局部变量m

	//fmt.Println(name, age, price, isOk, m)

	//5. 匿名变量  在使用多重赋值时，如果想要忽略某个值 可以使用 _
	x, _ := foo(10, "abc")
	_, y := foo(10, "abc")
	fmt.Println(x, y)


}
func foo(x int, y string) (int, string) {
	return x, y
}

