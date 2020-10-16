package main

import "fmt"

//arguments 可变参数
//args... 一定是最后一个形参 不能放前面或中间

func sum1(args...int) (sum int) { // 0 - 多个参数
	for i := 0; i < len(args); i++ {
		sum += args[i]
	}
	return
}
func sum2(n1 int, args...int) (sum int) {
	sum = n1
	for _, val := range args {
		sum += val
	}
	return
}

func main() {
	s1 := sum1(1,3,4,5,6,7)
	s2 := sum2(1,3,4,5,6,7)
	fmt.Println("s1=",s1, " s2=",s2)
}
