package main

import "fmt"

func main() {
	intChan := make(chan int, 3)
	intChan<- 3
	close(intChan) //关闭后无法继续写入，但可以读取

	num := <- intChan
	fmt.Print(num)
}
