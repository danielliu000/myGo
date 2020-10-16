package main

import (
	"fmt"
	"sync"
)

//协程统计1- n(200000)中的素数
//思路: 1. 创建一个方法setNum 开启一个协程存放1- n的数字, 放入intChan
//      2. 创建一个getPrime方法, 开启多个协程 统计素数 放入 primeChan
//      3. 创建一个printPrime方法,开启一个协程从primeChan中取素数并打印
//      4. 由于时多协程通信 需要sync.waitGroup 告知主线程何时结束
//      5. intChan primChan完成后需要关闭
//      6. 但primeChan由于时多协程,需要确定何时多协程全部完成后 channel关闭,
//     		 否则printPrime()取数是会造成deadlock.
//			7. 这时引入exitChan, 每执行一个primeChan协程 写入一个标志位,
//	    	 告知协程完毕.从exitChan中取值, 全部取完后, 即primeChan全部协程执行完毕,
//         这时关闭primeChan,exitChan
var wg = sync.WaitGroup{}


func setNum(intChan chan int, num int) {
	for i := 2; i < num; i++ {
		intChan <- i
	}
	close(intChan)
	wg.Done()
}
func getPrime(intChan, primeChan chan int, exitChan chan bool) {
	for v := range intChan {
		flag := true
		for i := 2; i < v; i++ {
			if v % i == 0 {
				flag = false
				break
			}
		}
		if flag {
			primeChan <- v
		}
	}
	exitChan <- true
	wg.Done()


}

func printPrime(primeChan chan int) {
	for v := range primeChan {
		fmt.Println(v)
	}
	wg.Done()
}


func main() {
	const PrimeRoutine = 20  //获取素数协程的个数
	num := 20000   // num内素数
	intChan := make(chan int, num)
	primeChan := make(chan int, 200)
	exitChan := make(chan bool, PrimeRoutine)

	//写入数字协程
	wg.Add(1)
	go setNum(intChan, num)

	//获取素数协程
	for i := 0; i < PrimeRoutine; i++ {
		wg.Add(1)
		go getPrime(intChan, primeChan, exitChan)
	}

	//打印素数协程
	wg.Add(1)
	go printPrime(primeChan)

	//退出primeChan协程
	wg.Add(1)
	go func(){
		for i := 0; i < PrimeRoutine; i++ {
			<- exitChan
		}
		close(primeChan)
		close(exitChan)
		wg.Done()
	}()


	wg.Wait() //等待所有协程执行完毕
	fmt.Println("程序执行完毕")
}
