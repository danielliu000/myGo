package main

import "fmt"

func main() {
	// 4种初始化数组方式：
	var numArr01 =  [3]int{1,2,3}
	fmt.Println("numArr-01",numArr01)

	var numArr02 [3]int = [3]int {4,5,6}
	fmt.Println("numArr-02",numArr02)

	var numArr03 = [...]int {7,8,9}
	fmt.Println("numArr-03",numArr03)

	var numArr04 = [...]int{1: 10, 2: 23,3: 34, 0: 3}
	fmt.Println("numArr-04",numArr04)
}
