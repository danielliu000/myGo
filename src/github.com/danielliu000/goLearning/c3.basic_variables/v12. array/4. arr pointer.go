package main

import "fmt"

func test01 (arr [3]int) {
	arr[0] = 99
	fmt.Println(arr)
}

func test02(arr *[3]int) {
	(*arr)[0] = 88
	fmt.Println(arr)
}

func main() {
	arr := [3]int{11,22,33}
	test01(arr) // 数组是值类型，不会改变原值
	fmt.Println(arr)
	test02(&arr) //指针会改变原值
	fmt.Println(arr)

}
