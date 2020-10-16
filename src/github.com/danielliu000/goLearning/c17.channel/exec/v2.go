package main

import (
	"fmt"
)
//练习
//1. 启动一个协程，将1-2000的数放入一个channel, 比如numChan
//2. 启动8个协程，从channel中取出数(n)，并计算1+...+n的值，并存放到resChan
//3. 最后8个协程协同完成工作后，再遍历resChan,显示结果 如： res[1] = 1... res[10] = 55..
//4. 注意： 考虑resChan chan int是否合适

func writeToChan(numChan chan int) {
	for i := 1; i <= 2000; i++ {
		numChan <- i
		fmt.Printf("wirteTochan : %v\n", i)
	}
	close(numChan)
}

func sumToChan(numChan chan int, resChan chan uint64) {
	for v := range numChan {
		var res uint64 = 0
		var i uint64
		for i = 1; i <= uint64(v); i++ {
			res += i
		}
		resChan <- res
		fmt.Printf("res[%v] = %v\n", v, res)
	}
	flag = true
}

var flag = false
func main() {

	numChan := make(chan int, 2000)
	resChan := make(chan uint64, 2000)

	for i := 0; i < 8; i++ {
		go sumToChan(numChan, resChan)
	}
	go writeToChan(numChan)
	for {
		if flag == true {
			break
		}
	}

}
