package main

import (
	"fmt"
	"strconv"
	"time"
)

func test() {
	for i := 1; i<= 100000; i++ {
		str := "hello world" + strconv.Itoa(123)
		fmt.Println(str)
	}
}
func main() {
	//1. 统计一个函数的执行时间
	startTime := time.Now()
	test()
	duration1 := time.Since(startTime)
	duration2 := time.Now().Sub(startTime)
	fmt.Println("执行test花了: ", duration1)
	fmt.Println("执行test花了: ", duration2)

}
