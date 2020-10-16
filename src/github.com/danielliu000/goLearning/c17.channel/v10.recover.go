package main

import (
	"fmt"
	"sync"
	"time"
)

//使用recover 解决协程panic

func sayHello() {
	for i := 0; i < 10; i++ {
		time.Sleep(time.Second)
		fmt.Println("hello, world")
	}
	wg3.Done()
}
func test() {
	//这里使用defer + recover  不会影响其他协程和主线程
	defer func() {
		if err := recover(); err != nil {
			fmt.Println("test()发生错误", err)
		}
	}()
	var myMap map[int]string // 没有make
	myMap[0] = "golang"  // error: entry in nil map
	wg3.Done()
}
var wg3 = sync.WaitGroup{}
func main() {

	wg3.Add(1)
	go sayHello()
	wg3.Add(1)
	go test()   //
	wg3.Wait()
}
