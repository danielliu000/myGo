package main

import (
	"fmt"
	"strconv"
	"sync"
	"time"
)
var wg = sync.WaitGroup{}  //解决主线程提前退出的情况
func test1() {
	for i := 0; i < 10; i++ {
		fmt.Println("Hello World " + strconv.Itoa(i))
		time.Sleep(time.Second * 2)
	}
	wg.Done() //协程计数器-1
}
func main() {
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go test1()
	}
	for i := 0; i < 10; i++ {
		fmt.Println("Hello Golang! " + strconv.Itoa(i))
		time.Sleep(time.Second)
	}
	wg.Wait() //等待协程退出
	fmt.Println("主程序退出")
}
