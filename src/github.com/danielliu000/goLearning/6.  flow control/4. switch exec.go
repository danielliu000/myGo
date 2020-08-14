package main

import (
	"fmt"
	"utils"
)

//1. 字符a,b c,d,e,f,g 表示周一到周日

func getDate() {

	var key byte
	fmt.Printf("input your char: ")
	_, _ = fmt.Scanf("%c", &key) //注意格式是字符输入 %c

	switch key {
	case 'a':
		fmt.Println("Monday")
	case 'b':
		fmt.Println("Tuesday")
	case 'c':
		fmt.Println("Wednesday")
	case 'd':
		fmt.Println("Thursday")
	case 'e':
		fmt.Println("Friday")
	case 'f':
		fmt.Println("Saturday")
	case 'g':
		fmt.Println("Sunday")
	default:
		fmt.Println("wrong input")

	}
}

//2. 根据输入的月份和年份，求该月的天数
// 1/3/5/7/8/10/12 --- 31
// 2-- 29/28
// 其他 30
func getDays(month byte, year uint16) byte {
	switch month {
	case 1, 3, 5, 7, 8, 10, 12:
		return 31
	case 2:
		if utils.IsLeapYear(year) {
			return 29
		} else {
			return 28
		}
	default:
		return 30
	}
}


//3. 根据公式 (height - 108）*  2 = weight， 可以有10斤左右浮动，观察测试者体重是否合适

func weightTest(height uint16, weight uint16) string {
	stdWeight := (height - 108) * 2

	switch  {
	case weight > stdWeight + 10:
		return "You are overweight!"
	case weight < stdWeight - 10:
		return "You are underweight!"
	default:
		return "You are fit!"
	}

}


func main() {
	//days :=getDays(2,2001)
	//	//fmt.Println(days)
	//
	//isFit := weightTest(177, 140)
	//fmt.Println(isFit)


}
