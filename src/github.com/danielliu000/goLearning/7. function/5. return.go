package main

import "fmt"

func calNums(x, y int) (sum, sub int) { //支持给返回值命名

	//定义不用关心顺序
	sub = x - y //sub返回值
	sum = x + y //sum返回值

	return  //无需再声明返回
}


func main() {
	sum, sub := calNums(3,6)  //接收按声明顺序

	fmt.Println("sum: ",sum, " sub: ", sub)
}
