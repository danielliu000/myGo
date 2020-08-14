package main

import "fmt"

func main() {
	arr := [5]int{1,2,3,4,5}
	slice := arr[1:3] //引用arr数组的下标1 -（3-1）号元素   // 【2，3】

	fmt.Println(slice)
	fmt.Println(len(slice))
	fmt.Println(cap(slice))// 容量不确定

	fmt.Printf("%p\n", &arr[1])  // &mySort[1] == &slice[0]
	fmt.Printf("%v", &slice[0]) //其实就是指向arr[1]的内存地址
}
