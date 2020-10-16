package main

import "fmt"

func main() {
	myChan := make(chan int, 100)
	for i := 0; i < 100; i++  {
		myChan <- i * 2
	}
	//遍历: 不能使用for，因为无法确定长度, 需要使用for range
	close(myChan)
	for v := range myChan { //注意： channel没有index
		fmt.Println(v) //如果channel不关闭就遍历，会产生死锁deadlock：取空了也停不下来

	}


}
