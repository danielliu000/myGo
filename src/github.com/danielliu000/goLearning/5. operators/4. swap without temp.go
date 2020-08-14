package main

import (
	"fmt"
)
//交换a, b 不使用中件变量
//方式一
func swapA(a int, b int)  {
	a, b = b, a
}

//方式二
func swapB(a int, b int) {
	  a = a + b
	  b = a - b
	  a = a - b
}
func main() {

	a := 3
	b := 5
	swapA(a,b)
	fmt.Println(a, b)
	swapB(a,b)
	fmt.Println(a, b)
}
