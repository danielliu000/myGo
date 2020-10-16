package main

import "fmt"

func main() {
	arr := []int {1,2,3,4,5}
	slice1 := make([]int, 10)
	slice2 := make([]int, 3) // 长度不够 也可以拷贝,只能拷贝前3个
	fmt.Println(slice1)

	len1 := copy(slice1, arr)  // 必须都是切片 才可以拷贝
	len2 := copy(slice2, arr)
	fmt.Println(slice1, len1) //  返回赋值元素的个数  5
	fmt.Println(slice2, len2) //   3


	slice1[0] = 10
	fmt.Println(slice1)
	fmt.Println(slice2)
	fmt.Println(arr)
}
