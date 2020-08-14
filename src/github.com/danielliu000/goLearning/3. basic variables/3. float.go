package main

import (
	"fmt"
	"math"
)

func main() {
	//float32  单精度 4字节  约6个十进制数的精度
	//float64  双精度 8字节  约15个十进制数的精度


	//当整数大于23bit能表达的范围时，float32的表示将出现误差
	var num1 float32 = 16777216  //
	fmt.Println(num1 == num1 + 1)

	for x := 0; x < 8; x++ {
		fmt.Printf("x = %d e^x = %8.3f\n", x, math.Exp(float64(x)))
	}

	//+inf-inf 正负无穷大 太大溢出的数字和除零的结果
	//NaN非数，一般用于表示无效的除法操作结果0/0或Sqrt(-1)
	var z float64
	fmt.Println(z, -z, 1/z, -1/z, z/z, math.Sqrt(-1)) // "0 -0 +Inf -Inf NaN NaN"

}

