package main

import "fmt"

func main() {
	//channel默认是双向的

	//send only
	var chan1 chan <- int
	chan1 = make(chan int)
	chan1 <- 20
	fmt.Println(chan1) //可以打印地址
	//但不能取值
	//num := <- chan1  //error


	//receive-only
	var chan2 <- chan int  // 无法send/写入 所以一般用于函数参数
	chan2 = make(chan int)
	num2 := <- chan2
	//chan2 <- 20  // error
	fmt.Println(num2)

}
