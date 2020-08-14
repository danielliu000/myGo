package main

import "fmt"

func test1(n1 *int) {
	*n1 = *n1 + 20  // 引用类型操作的就是main 中的n， 外部n值会改变
	fmt.Println("n=", *n1)

}
func main() {
	n := 10
	test1(&n)   //引用类型
	fmt.Println("n=", n)
}
