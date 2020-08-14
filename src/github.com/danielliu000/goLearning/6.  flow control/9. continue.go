package main

import "fmt"

func main() {
	//for i := 0; i < 4; i++ {
	//	for j := 0; j < 10; j++ {
	//		if j == 2 {
	//			continue // 不执行循环体下面的代码，直接执行下一次循环
	//		}
	//		fmt.Println("j=", j)
	//	}
	//}
	here:
	for i := 0; i < 2; i++ {
		for j := 1; j < 4; j++ {
			if j == 2 {
				continue here
			}
			fmt.Println("i=",i, " j=",j)
		}
	}
}
