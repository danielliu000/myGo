package main

import "fmt"

//goroutine与channel协同

//开启writeData协程，向channel写入50个整数
//开启readData协程，读取channel数据
//注意： channel为引用类型，因此二个协程会操作同一个channel
//要保证二个协程工作时，主线程不会提前退出
func writeData1(intChan chan int) {
	for i := 1; i <= 50; i++ {
		intChan<- i
		fmt.Printf("writeData 写数据: %v\n", i)
	}
	close(intChan)
}
func readData1(intChan chan int, exitChan chan bool)  {
	for {
		v, ok := <-intChan
		if !ok {
			break
		}
		fmt.Printf("readData 读到数据= %v\n", v)
	}
	exitChan<- true
	close(exitChan)
}
func main() {

	intChan := make(chan int, 50)
	exitChan := make(chan bool, 1)

	go readData1(intChan, exitChan)
	go writeData1(intChan)

	for {
		_, ok := <-exitChan
		if !ok {
			break
		}
	}
}
//