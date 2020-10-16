package main

import "fmt"

func test() {
	//使用 defer + panic的方式捕获处理异常
	defer func() {
		 //recover 可以捕获异常
		if err := recover(); err != nil {
			fmt.Println("err=", err)
		}
	}()
	num1 := 10
	num2 := 0
	fmt.Println(num1 / num2)
}

func main() {
	test()
	fmt.Println("continue conducting codes...")
}
