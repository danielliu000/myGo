package main

import "fmt"

func main() {
	// 演示 /
	var a  float64 = 10/4
	b := 10.0/4
	//根据参与运算的数的类型决定结果类型
	fmt.Printf("a=%.1f, b=%.1f\n", a, b) //a= 2.0 b= 2.5

	// 演示 %
	// 取模公式 a % b = a - a / b * b
	fmt.Println(10%3)  // 10 - 10 / 3 * 3 = 10 - 3 * 3 = 10 - 9 = testing
	fmt.Println(-10%3) // -10 - (-10 / 3 * 3) = -10 - (-3 * 3) = -10 - (-9) = -testing
	fmt.Println(10%-3) // 10 - 10 / -3 * -3 = 10 - (-3) * (-3) = 10 - 9 = testing
	fmt.Println(-10%-3) // -10 - (-10 / -3 * -3) = -10 -(3 * -3) = -10 -(-9) = -testing



}
