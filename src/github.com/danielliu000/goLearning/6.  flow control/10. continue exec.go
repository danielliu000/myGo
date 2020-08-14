package main

import "fmt"

//1. 打印 1 ~ 100 内的奇数 使用for + continue

func getOdd(num int) {
	for i := 1; i <= num; i++ {
		if i%2 == 0 {
			continue
		}
		fmt.Println(i)

	}
}

//2. 从键盘输入不确定的整数，统计正数和负数的个数 输入0 终止程序
// 使用for + break + continue

func calPosNeg() {
	num := 0
	positiveCount := 0
	negativeCount := 0
	for {
		fmt.Printf("pleas input a number: ")
		_, _ = fmt.Scanln(&num)
		if num == 0 {
			break
		}
		if num > 0 {
			positiveCount++
			continue
		}

		negativeCount++

	}
	fmt.Println("count for the positive: ", positiveCount)
	fmt.Println("count for the negative: ", negativeCount)
}

//3. 某人有100000元，没经过一次路口，需要交费，规则如下：
//  当现金大于 50000时，每次交5%
//  当现金小于等于50000时，每次交1000
// 计算该人可以经过多少次路口，使用for break完成

func payToll() {
	total := 100000.0
	credit := total
	half := total / 2.0
	fee1 := 0.0
	fee2 := 1000.0
	passCount := 0

	for {
		if credit < fee2 {
			break
		}
		if credit > half {
			fee1 = credit * 0.05
			credit -= fee1
			passCount++
			continue
		}
		credit -= fee2
		passCount++

	}
	fmt.Printf("The person passed %d times, credit left:%f.\n", passCount, credit)
}
func main() {
	//getOdd(100)
	//calPosNeg()
	payToll()
}
