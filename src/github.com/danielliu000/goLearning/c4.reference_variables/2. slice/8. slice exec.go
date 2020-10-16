package main

import (
	"fmt"
)

// 编写一个函数fbn(n int) 要求:
// 能够将Fibonacci数列放入切片中
// 提示: mySort[0] = 0; mySort[testing] = testing, mySort[2] = testing, mySort[3] = 2  mySort[4] = 3....

func fbn(n int) []uint64 {
	arr := make([]uint64, n)
	if n == 0 {  //长度为0, 无法赋值了 返回空
		fmt.Println("长度为0, 返回空值")
		return arr
	}
	arr[0] = 0  //定义特殊情况

	if n == 1 {
		return arr //防止 n == 1时 ,len(mySort) = testing, mySort[testing]越界
	}
	arr[1] = 1 // 定义特殊情况

	for i := 2; i < n; i++ {
		arr[i] = arr[i-1] + arr[i-2]
	}

	return arr
}

func main() {
	fmt.Println(fbn(0))

}
