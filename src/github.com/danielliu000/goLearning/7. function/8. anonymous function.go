package main

import "fmt"

var ( //定义全局匿名函数
	func1 = func() {
		fmt.Println("I am a global anonymous function: func1 ")
	}
	func2 = func() {
		fmt.Println("I am a global anonymous function: func2 ")
	}
)

func main() {

	//在main中定义了匿名函数
	res1 := func(a, b int) int {
		return a + b
	}(1, 2) //直接传参 res1 类型为整型
	fmt.Printf("res1 type: %T\n", res1)
	fmt.Println(res1)

	res2 := func(a, b int) int {
		return a - b
	} //不传参，res2此时为 func(int, int) int 类型
	fmt.Printf("res1 type: %T\n", res2)
	fmt.Println(res2(2, 4))


	//执行全局匿名函数
	func1()
	func2()
}
