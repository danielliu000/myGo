package main

import "fmt"

func main() {
	intChan := make(chan int, 10)
	for i := 0; i < 10; i++ {
		intChan <- i
	}

	stringChan := make(chan string, 5)
	for i := 0; i < 5; i++ {
		stringChan <- "hello" + fmt.Sprintf("%d", i)
	}

	//传统方法遍历管道时， 如果不关闭会阻塞导致deadlock
	//但有时又不确定何时关闭管道
	//这时， 可以使用select 方法 依case 来取 直到取不出来

	Loop:
	for {
		select {
		case v := <-intChan:
			fmt.Printf("intChan数据：%d\n", v)
		case v := <-stringChan:
			fmt.Printf("stringChan数据：%s\n", v)
		default:
			fmt.Printf("都取不到，不玩了\n")
			break Loop
		}
	}

}
