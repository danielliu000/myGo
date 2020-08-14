package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	//basic: show date and time
	now := time.Now()
	fmt.Printf("now = %v, type = %T\n" , now, now)
	//-------------------------------------//

	//format time and date
	//方式一 : Printf 或者SPrintf
	fmt.Printf("%02d-%02d-%02d, %02d:%02d:%02d\n",
		now.Year(), now.Month(), now.Day(),
		now.Hour(), now.Minute(),now.Second())
	//方式二： 使用 now.Format() layout: "2006/01/02 15:04:05" 数字不能改:(
	fmt.Println(now.Format("2006/01/02 15:04:05"))
	//------------------------------------------//

	//时间常量 结合Sleep
	for i:=1; i <= 5; i++ {
		fmt.Println("i=",i)
		//time.Sleep(2 * time.Second)   // 2s
		time.Sleep(100 * time.Millisecond) //0.1s
		//time.Sleep(2 * time.Minute)  //2 min
		//time.Sleep(2 * time.Second)
	}
	//----------------------------------------------//

	//Unix UnixNano 随机数 使用
	fmt.Println("Unix时间戳:", now.Unix())
	fmt.Println("UnixNano时间戳:", now.UnixNano())

	//将时间戳设置成种子数
	rand.Seed(time.Now().UnixNano())
	//生成10个0-99之间的随机数
	for i:=0;i<10;i++{
		fmt.Println(rand.Intn(100))
	}
}


