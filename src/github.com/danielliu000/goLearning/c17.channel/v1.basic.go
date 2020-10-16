package main

import "fmt"

func main() {
	//创建一个可以存放3个int的channel
	var intChan chan int
	intChan = make(chan int, 3)
	fmt.Printf("intChan'value: %v\n", intChan)  //channel为引用类型，intChan指向一个地址
	fmt.Printf("intChan'value: %v\n", &intChan) //intChan本身的地址

	//向管道写数据
	intChan <- 100
	num2 := 200
	intChan <- num2

	//查看管道长度容量=>容量不会扩容，这点与切片不同
	//但可以从管道中取内容，取出后长度自动变小，从而可以继续放入
	fmt.Printf("channel len: %v, channel cap: %v\n", len(intChan), cap(intChan))

	num11 := <-intChan
	num22 := <-intChan
	fmt.Printf("num2= %v\n", num11)
	fmt.Printf("num2= %v\n", num22)
	fmt.Printf("channel len: %v, channel cap: %v\n",
		len(intChan), cap(intChan)) // len : 0 cap: 3
		//没有使用协程时，从空channel中取值会报错
		//num3 := <- intChan
		//fmt.Println(num3) // deadlock错误
}