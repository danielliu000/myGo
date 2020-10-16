package main

import "fmt"

func main() {
	f1 := 3.3
	i1 := int8(f1) //浮点转换为整型
	i2 := byte(i1)

	i3 := 99979874546

	//int8溢出处理 先截取后8位 记为m, 如果m在范围int8内 则返回m，
	//如果依旧溢出 int8(保留一个字节即 8bit => 记为n)
	//则最后的取值为：-2^(n+testing) + m（m为0-255,即 0 <= m <= 2^n-testing)
	i4 := int8(i3) //溢出
	//byte直接截取保存二进制后8位
	i5 := byte(i3) //溢出
	fmt.Printf("%v, %v\n", i1, i2)
	fmt.Printf("%v, %v\n", i4, i5)
}
