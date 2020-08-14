package main

import (
	"fmt"
	"strconv"
)

func main() {
	// 基本类型转为string
	//
	//
	//方式一： 使用 fmt.Sprintf()
	var num1 int = 99
	var num2 float64 = 23.456
	var b bool = true
	var myChar byte = 'h'
	var str string

	str = fmt.Sprintf("%d", num1)
	fmt.Printf("str Type：%T, str value: %q\n", str, str)

	str = fmt.Sprintf("%.3f", num2) //保留精度
	fmt.Printf("str Type：%T, str value: %q\n", str, str)

	str = fmt.Sprintf("%t", b)
	fmt.Printf("str Type：%T, str Value: %q\n", str, str)

	str = fmt.Sprintf("%c", myChar)
	fmt.Printf("str Type：%T, str Value: %q\n", str, str)

	//方式二： 使用Package strconv
	fmt.Println("-------------------------------")
	fmt.Println("Use Package strconv to convert:")

	str = strconv.FormatInt(int64(num1), 10)
	fmt.Printf("str Type：%T, str value: %q\n", str, str)
	// strconv.Itoa int to string 方法
	str = strconv.Itoa(int(num1))
	fmt.Printf("str Type：%T, str Value: %q\n", str, str)

	str = fmt.Sprintf("%.3f", num2) //保留精度
	fmt.Printf("str Type：%T, str value: %q\n", str, str)

	str = strconv.FormatFloat(num2, 'f', 10, 64)
	// 'f' 格式 dddd.ddd
	// 'prec 10' 小数精度 保留10位小数
	// 'bitSize 64  小数float64
	fmt.Printf("str Type：%T, str Value: %q\n", str, str)

	str = strconv.FormatBool(b)
	fmt.Printf("str Type：%T, str Value: %q\n", str, str)


	// itoa int to string 方法
	str = strconv.Itoa(num1)
	fmt.Printf("str Type：%T, str Value: %q\n", str, str)




}
