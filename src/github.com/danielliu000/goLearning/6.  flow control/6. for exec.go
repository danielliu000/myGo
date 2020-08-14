package main

import (
	"fmt"


)

//1. 打印1-100 被9整除的个数及其总和
func sumCount(maxNum uint, divisor uint) (uint, uint) {
	var count uint = 0
	var sum uint = 0

	var i uint = 1
	for ; i <= maxNum; i++ {
		if i%divisor == 0 {
			count++
			sum += i
		}
	}
	return count, sum
}

//2. 统计三个班的各自平均分 ，总平均分，以及及格总人数，每班有5个学生，要求成绩键盘输入
func getAvg(classNum int, studentNum int) {
	var (
		classAvg  float32 = 0
		sumClass  float32 = 0
		sumGrade  float32 = 0
		gradeAvg  float32 = 0
		score     float32 = 0
		passCount uint    = 0
	)
	for i := 1; i <= classNum; i++ {
		sumClass = 0
		for j := 1; j <= studentNum; j++ {
			fmt.Printf("输入班级[%d]，学生（%d）的成绩: ", i, j)
			_, _ = fmt.Scanln(&score)
			if score >= 60 {
				passCount++
			}
			sumClass += score
		}
		classAvg = sumClass / float32(studentNum)
		fmt.Printf("班级[%d]的平均分是:%.2f \n", i, classAvg)

		sumGrade += classAvg
	}
	gradeAvg = sumGrade / float32(classNum)
	fmt.Printf("年级平均分是：%.2f, 及格人数为: %d", gradeAvg, passCount)

}

//3. 打印金字塔经典案例
func getPyramid(layer int, isSolid bool) {
	for i := 1; i <= layer; i++ { //层数
		for k := 1; k <= layer-i; k++ {// 打印空格,个数为: 总层数-当前层数
			fmt.Printf(" ")
		}
		for j := 1; j <= 2*i-1; j++ { // 每层*的个数 2 * 当前层 - 1
			if isSolid { // 打印实心金字塔
				fmt.Printf("*")
			} else { // 打印空心金字塔
				if j == 1 || j == 2 * i - 1 || i == layer {
					fmt.Printf("*")
				} else {
					fmt.Printf(" ")
				}
			}
		}
		fmt.Println()
	}
}

//4. 打印九九乘法表
func getMultiTable (num int) {
	for x := 1 ; x <= num; x++ {
		fmt.Println()
		for y :=1; y <= x; y++ {
			fmt.Printf("%d x %d = %d\t", y, x, x * y)
		}

	}
}

func main() {
	//fmt.Println(sumCount(100, 9))
	//getAvg(3, 3)
	//getPyramid(10, false)
	getMultiTable(9)
}
