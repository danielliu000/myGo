package main

import "fmt"

func main() {
		//定义声明二维数组

	var arr [4][6]int
	fmt.Println(arr)
	//赋初值
	arr[1][2] = 1
	arr[2][1] = 2
	arr[2][3] = 3
	for i := 0; i < 4; i++ {
		for j:= 0; j < 6; j++ {
			fmt.Printf("%v ",arr[i][j])
		}
			fmt.Println()
	}

	var arr2 [2][3]int = [2][3]int {{1,2,3}, {4,5,6}}


	fmt.Println(arr2)

}
