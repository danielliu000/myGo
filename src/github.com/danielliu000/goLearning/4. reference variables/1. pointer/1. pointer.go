package main

import "fmt"

func main() {
	// 获取变量地址 使用 & + Variables Name
	i := 10
	fmt.Println("i' address: ", &i)
	var ptr *int = &i
	//1. ptr 是一个指针变量
	//2. ptr 的类型是 *int (指针类型 有 int float bool string array struct
	//3. ptr 本身的值为&i
	fmt.Println("ptr's value:", ptr)  // ptr指针存储的内容 ：ptr 即 i的地址
	fmt.Println("ptr's own address in memory:", &ptr) //ptr存储内容的内存地址: &ptr
	fmt.Println("ptr's pointing value: ", *ptr)   //ptr指向空间的内容 *ptr 即 i的值

	num1 := 100
	fmt.Printf("num1 type: %T, num1 value: %v, num1 addr: %v\n", num1, num1, num1)

	num2 := new(int)
	fmt.Printf("num2 type: %T " +
		"num2 value: %v " +
		"num2 addr: %v " +
		"num2 pointing addr: %v\n",
		num2, *num2, &num2, num2)


}
