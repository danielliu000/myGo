package main

import "fmt"

func main() {
	myMap := map[int]int{
		1: 200,
		2: 400,
		10: 33,
		5: 33,
		3:20,
	}
	fmt.Println(myMap) //v1.14.2后 默认自动排序了
}
