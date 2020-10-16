package main

import (
	"bufio"
	"errors"
	"fmt"
	"math/rand"
	"os"
	"strconv"
	"strings"
	"time"
	"utils"
)

//Q1. 循环打印输入的月份天数, 使用continue
// 要求判断输入是否错误  分别提示输入的年/月/日不正确
//请输入年:
//请输入月:
//请输入日:
//您输入的日期是 xxxx年xx月xx日

var (
	input         = ""
	err     error = nil
	r             = bufio.NewReader(os.Stdin)
	isPrime       = false
)

func getDate() (string, string, string) {
	var (
		y = 0
		m = 0
		d = 0
	)
	for {
		y, err = checkInput("请输入年:", y)
		if err != nil {
			fmt.Println(errors.New("输入的年不正确"))
			continue
		}

		m, err = checkInput("请输入月:", m)
		if err != nil || m > 12 || m < 1 {
			fmt.Println(errors.New("输入的月不正确"))
			continue
		}

		d, err = checkInput("请输入日:", d)
		if err != nil || !isValidDay(y, m, d) {
			fmt.Println(errors.New("输入的日不正确"))
			continue
		}
		//格式化日期 补零
		//year := strconv.Itoa(y)
		//month := strconv.Itoa(m)
		//day := strconv.Itoa(d)
		//if d < 10 {
		//	day = "0" + day
		//}
		//if m < 10 {
		//	month = "0" + month
		//}
		//fmt.Printf("您输入的日期是: 公元%v年%v月%v日\n", year, month, day)
		//return year, month, day
		year := fmt.Sprintf("%v", y)
		month := fmt.Sprintf("%02v", m)
		day := fmt.Sprintf("%02v", d)
		fmt.Printf("您输入的日期是: 公元%v年%v月%v日\n", year, month, day)
		return year, month, day
	}

}
func checkInput(inputLabel string, inputConv int) (int, error) {
	fmt.Printf(inputLabel)
	input, err = r.ReadString('\n')
	return strconv.Atoi(strings.ReplaceAll(input, "\n", ""))

}
func isValidDay(y int, m int, d int) bool {
	isLeap := utils.IsLeapYear(y)
	if d > 31 || d < 1 {
		return false
	}
	if !isLeap && m == 2 && d > 28 {
		return false
	}
	switch m {
	case 4, 6, 9, 11:
		if d == 31 {
			return false
		}
	}
	return true

}

//Q2. 随机猜数游戏
//    随机生成一个1-100的整数, 有10次机会,
// 第 testing 次猜中 提示"你是天才"
// 第 2-3 次猜中 提示 "你很聪明,赶上我了"
// 第 4-9 次猜中 提示"一般般"
// 第 10 次猜中 提示"可算猜对了"
// 一次没猜中 提示"说你什么好呢"
func gussNums() {
	num := 0
	rand.Seed(time.Now().UnixNano())
	randomNum := rand.Intn(100)
	fmt.Println(randomNum)

	count := 1
Loop:
	for ; count <= 10; count++ {
		num, err = checkInput("数字是: ", num)

		if err != nil {
			fmt.Println(errors.New("输入出错了,浪费一次机会"))
			continue
		}
		if err == nil && randomNum == num {
			switch count {
			case 1:
				fmt.Println("你是天才")
				break Loop
			case 2, 3:
				fmt.Println("你很聪明,赶上我了")
				break Loop
			case 4, 5, 6, 7, 8, 9:
				fmt.Println("一般般")
				break Loop
			case 10:
				fmt.Println("可算猜对了")
			}
		}

	}
	if count > 10 {
		fmt.Println("说你什么好呢")
	}

}

//Q3 编写一个函数, 输出100内所有素数, 每行显示5个 并求和
func getPrimeNum(num uint) {
	var i uint = 2
	count := 0
	for ; i < num; i++ {
		isPrime = true
		var j uint = 2
		for ; j < i; j++ {
			if i%j == 0 {
				isPrime = false
				break
			}
		}
		if isPrime {
			count++
			fmt.Printf("%v ", i)
			if count == 5 {
				count = 0
				fmt.Printf("\n")
			}
		}
	}

}

//Q4 从1990-01-01起,  "三天打鱼两天筛网" 判断以后某一天是打鱼还是晒网
func isFishingDate() string {
	y, m, d := getDate()

	if year, _ := strconv.Atoi(y); year < 1990 {
		return "输入的年不正确，请输入1990年之后的日期"
	}
	timeFormatTpl := "2006-01-02"
	startDate := "1990-01-01"
	endDate := y + "-" + m + "-" + d

	s, _ := time.Parse(timeFormatTpl, startDate) // Parse 第一个参数是layout 即日期格式
	e, _ := time.Parse(timeFormatTpl, endDate)
	duration := e.Sub(s).Hours() / 24

	if int(duration) % 5 < 3 {
		fmt.Println(e,duration)
		fmt.Println(endDate)
		return "今天打鱼"
	} else {
		return "今天筛网"
	}
}

//Q5 编写一个计算器 计算 + - * / 要求有界面, 可以根据选择进行计算
func calculator() {
	num1 := 10
	num2 := 5
	exit := 0
	add := 1
	sub := 2
	mul := 3
	div := 4
	input := 0
	fmt.Println("-----简易计算器-----")
	fmt.Printf("%v. +\n", add)
	fmt.Printf("%v. -\n", sub)
	fmt.Printf("%v. *\n", mul)
	fmt.Printf("%v. /\n", div)
	fmt.Printf("%v. exit\n", exit)
	fmt.Println()

Loop:
	for {
		//fmt.Printf("请输入选项：")
		//_, _ = fmt.Scanln(&input)
		input, err = checkInput("请输入选项：", input)
		switch {

		case input == 1:
			fmt.Println(num1 + num2)

		case input == 2:
			fmt.Println(num1 - num2)

		case input == 3:
			fmt.Println(num1 * num2)

		case input == 4:
			fmt.Println(num1 / num2)
		case input == 0 && err == nil:
			fmt.Println("see you")
			break Loop
		default:
			fmt.Println("输入错误，重新输入")
		}
	}

}

//Q6 输出小写a-z以及大写Z-A
func getLetters() {
	upperLetter := 'Z'
	lowerLetter := 'a'
	fmt.Println("Letters:")
	for i := 0; i < 26; i++ {
		fmt.Printf("小写字母: %c --- 大写字母: %c\n", lowerLetter, upperLetter)
		lowerLetter++
		upperLetter--
	}

}
func main() {
	//getDate()
	//gussNums()
	//getPrimeNum(100)
	fmt.Println(isFishingDate())
	//calculator()
	//getLetters()
}
