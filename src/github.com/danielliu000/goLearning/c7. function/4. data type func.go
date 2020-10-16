package main

import "fmt"

func getSum(a int, b int) int {
	return a + b
}

func applySum(sumFunc func(int, int) int, x, y int ) int {
	return sumFunc(x, y)
}

func main() {
	//函数是一个数据类型：为 func (arguments types... ) return type
	fmt.Printf("getSum type: %T\n",getSum)

	res := applySum(getSum, 40,30) // 参数可以时一个函数
	fmt.Println(res)

}
