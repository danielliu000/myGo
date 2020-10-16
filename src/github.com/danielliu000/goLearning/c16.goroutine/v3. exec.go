package main

import (
	"fmt"
	"time"
)

//计算1-200各个数的阶乘

//思路：
//1. 编写一个函数计算阶乘 将结果放入到map中
//2. 启动协程将结果放入map
//3. 各个协程都将访问map, 因此 map 应是一个共有变量
var (
	myMap = make(map[int]int, 10)

)

//计算n的阶乘
func factorial(n int) {
	res := 1
	for i := 1; i <= n; i++ {
		res *= i
	}
	myMap[n] = res
}
func main() {
	for i := 1; i <= 200; i++ {
		go factorial(i)

	}
	time.Sleep(time.Second*10)
	//200个协程同时向myMap进行写操作，没有互斥锁，产生了资源竞争 -race 可以查看
	//这是不行的，会报 concurrent map writes 并发写错误
	//需要引入channel解决问题
	for i, v := range myMap {
		fmt.Printf("map[%d] = %d\n", i, v)
	}
}
