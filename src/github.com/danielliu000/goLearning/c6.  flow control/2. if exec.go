package main

import (
	"fmt"
	"math"
	"utils"
)

//testing. 定义2个变量int32, 判断二者的和是否能被3 和 5 整除， 给出提示信息

func getInfo(a int32, b int32) {
	sum := a + b
	if sum%3 == 0 && sum%5 == 0 {
		fmt.Println("a,b之和可以被3和5整除")
	} else if sum%3 == 0 && sum%5 != 0 {
		fmt.Println("a,b之和只能被3整除")
	} else if sum%3 != 0 && sum%5 == 0 {
		fmt.Println("a,b之和只能被5整除")
	} else {
		fmt.Println("a,b之和不能被3或5整除")
	}
}
func userInput() (int32, int32) {
	var a int32
	var b int32
	fmt.Printf("输入a的值: ")
	_, _ = fmt.Scanln(&a)
	fmt.Printf("输入b的值: ")
	_, _ = fmt.Scanln(&b)
	return a, b

}

//2. 判断是否为闰年
func isLeapYear(year int16) bool {
	if year%400 == 0 || (year%4 == 0 && year%100 != 0) {
		return true
	} else {
		return false
	}
}

//3. ax^2 + bx + c = 0 求 x
func getX(a float64, b float64, c float64) {
	g := b*b - 4*a*c
	if g > 0 {
		x1 := (-b + math.Sqrt(g)) / (2 * a)
		x2 := (-b - math.Sqrt(g)) / (2 * a)
		fmt.Printf("x1=%f, x2=%f\n", x1, x2)
	} else if g == 0 {
		x := (-b + math.Sqrt(g)) / (2 * a)
		fmt.Printf("x=%f\n", x)
	} else {
		fmt.Println("无解")
	}
}

//4. 出票系统:
//4-10月旺季， 成人(18-60): 60 儿童（<18): 30  老人（>60) :20
//淡季: 成人: 40, 其他: 20
func getTicketPrice() {
	var (
		tp1   byte = 60
		tp2   byte = 40
		month byte
		age   byte
	)
	fmt.Printf("输入月份: ")
	_, _ = fmt.Scanln(&month)

	fmt.Printf("输入年龄: ")
	_, _ = fmt.Scanln(&age)

	if month >= 4 && month <= 10 {
		if age >= 18 && age < 60 {
			fmt.Println("旺季成人票，票价为: ", tp1)
		} else if age < 18 {
			fmt.Println("旺季儿童票，票价为: ", tp1/2)
		} else {
			fmt.Println("旺季老年票，票价为: ", tp1/3)
		}
	} else {
		if age >= 18 && age < 60 {
			fmt.Println(" 淡季成人票，票价为: ", tp2)
		} else {
			fmt.Println(" 淡季老人儿童票，票价为: ", tp2/2)
		}

	}

}

//5. 判断一个数是否为水仙花数
func isNarcissisticNumber(num int) bool {
	if num >= 1000 || num <= 100 {
		return false
	} else {
		h := num / 100
		t := (num - h*100) / 10
		d := num - t*10 - h*100
		rs :=
			math.Pow(float64(h), 3) +
				math.Pow(float64(t), 3) +
				math.Pow(float64(d), 3)

		if num == int(rs) {
			return true
		} else {
			return false
		}

	}

}

//6. 求100 - 1000 的所有水仙花数
func getNarcissisticNumber() {
	for i := 100; i < 1000; i++ {
		if utils.IsNarcissisticNumber(i) {
			fmt.Printf("%v ", i)
		}
	}
}

//7. 下面代码输出结果是：
func getResult() {
	m, n := 0, 3
	if m > 0 {
		if n > 2 {
			fmt.Println("A")
		} else {
			fmt.Println("B")
		}
	} else {
		fmt.Println("print nothing")
	}
}

//8. 三个整数排序，从大到小输出
func sortNums(x int, y int, z int) {
	if y > x {
		x, y = y, x
	}
	if z > x {
		z, x = x, z
	}
	if z > y {
		z, y = y, z
	}
	fmt.Printf("%d, %d, %d", x, y, z)
}

func main() {
	//getInfo(userInput())
	//fmt.Println(isLeapYear(2020))
	//getX(testing,2,testing)
	//getTicketPrice()

	//var b bool = isNarcissisticNumber(200)
	//fmt.Println(b)
	//getNarcissisticNumber()
	//getResult()

}
