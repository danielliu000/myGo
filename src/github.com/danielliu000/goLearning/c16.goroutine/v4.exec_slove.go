package main

import (
	"fmt"
	"sync"
	"time"
)

//计算1-200各个数的阶乘
//引入 互斥锁 解决问题
var (
	myMap2 = make(map[int]uint64, 10)
	//引入互斥锁
	lock sync.Mutex // mutex是结构体 内部有两个方法Lock 和 Unlock
)

//计算n的阶乘
func factorial2(n int) {
	var res uint64 =  1
	for i := 1; i <= n; i++ {
		res *= uint64(i)
	}
	lock.Lock() //加锁
	myMap2[n] = res
	lock.Unlock() //解锁
}
func main() {
	for i := 1; i <= 200; i++ {  //阶乘数太大，数据类型其实写不下，会变成0
		go factorial2(i)
	}
	// 需要引入channel主线程解决问题
	time.Sleep(time.Second*10) // 协程结束前，主线程不能先结束 强制等待10s

	//200个协程同时向myMap进行写操作，没有互斥锁，产生了资源竞争 -race 可以查看
	//这是不行的，会报 concurrent map writes 并发写错误

	for i, v := range myMap2 {
		lock.Lock()
		fmt.Printf("map[%d] = %d\n", i, v)
		lock.Unlock()
	}
}
