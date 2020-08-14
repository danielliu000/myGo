package main

import (
	"fmt"
	"strconv"
)

func main() {
	//string to basic type

	//using package strconv
	// 注意 Parse... 方法有二个返回值 value,  err
 	// strconv.ParseInt()
	// strconv.ParseFloat()
	// strconv.ParseBool()



	str1 := "99"
	str2 := "23.456"
	str3 := "true"
	str4 := "97"
	str5 := "hello"

	num1, _ := strconv.ParseInt(str1, 10, 64) // ignore err argument
	fmt.Printf("num1 Type：%T, num1 value: %d\n", num1, num1)

	num2, _ := strconv.ParseFloat(str2, 64) // ignore err argument
	fmt.Printf("num2 Type：%T, num2 value: %.3f\n", num2, num2)

	b, _ := strconv.ParseBool(str3)
	fmt.Printf("b Type：%T, b value: %t\n", b, b)

	c, _ := strconv.Atoi(str4)   // "97" 数字是可以成功转换的
	fmt.Printf("c Type：%T, c value: %c\n", c, c)

	e, _ := strconv.Atoi(str5) // "hello"无法转， 虽转为int 但值为默认值
	fmt.Printf("e Type：%T, e value: %v\n", e, e)


}
