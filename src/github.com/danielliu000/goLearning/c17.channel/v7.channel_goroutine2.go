package main

import (
	"fmt"
	"sync"
)

//goroutine与channel协同

//开启writeData协程，向channel写入50个整数
//开启readData协程，读取channel数据
//注意： channel为引用类型，因此二个协程会操作同一个channel
//要保证二个协程工作时，主线程不会提前退出
func writeData2(intChan chan int) {
	for i := 1; i <= 50; i++ {
		intChan<- i
		fmt.Printf("writeData 写数据: %v\n", i)
	}
	close(intChan)
	wg1.Done()
}
func readData2(intChan chan int)  {
	for {
		v, ok := <-intChan
		if !ok {
			break
		}
		fmt.Printf("readData 读到数据= %v\n", v)
	}
	wg1.Done()
}

var wg1 sync.WaitGroup
func main() {

	intChan := make(chan int, 50)
	wg1.Add(1)
  go readData2(intChan)
	wg1.Add(1)
	go writeData2(intChan)
	wg1.Wait()

}
