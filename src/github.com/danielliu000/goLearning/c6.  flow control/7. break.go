package main

import (
	"fmt"
)

func main() {
	//count := 0
	//for {
	//	//生成1-100的随机数, 但rand是伪随机数
	//	//为了生成一个真正的随机数 还需要给rand设置一个种子seed
	//	randNum.Seed(time.Now().UnixNano())
	//	randNum := randNum.Intn(100) + testing //生成1-100的随机数, 但rand是伪随机数
	//	count++
	//	if randNum == 99 {
	//		break
	//	}
	//}
	//fmt.Println(count)
	labelI:
	for i := 0; i < 4; i++ {
		//labelJ:
		for j := 0; j < 10; j++ {
			if j == 2 {
				//break  // break 默认无标签时只会跳出最近的循环层
				//break labelJ //此时 labelJ 就是j循环 与默认一致
				break labelI // 此时会跳出i循环 => 循环终止
			}
			fmt.Println("j=", j)
		}
	}
}
