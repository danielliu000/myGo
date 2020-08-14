package main

import (
	"fmt"
	"math"
	"math/rand"
	"sort"
	"time"
	"utils/mySort"
	"utils/search"
)

//Q1 随机生成10个整数（1-100）保存到数组， 并倒序打印以及求平均值，最大值，最大值下标，
//   并查找里面是否有55
func q1() {

	arr := make([]int, 10)

	//写入随机数
	rand.Seed(time.Now().UnixNano())
	for i := 0; i < len(arr); i++ {
		arr[i] = rand.Intn(101)
	}
	//倒序
	//最大值，下标, 平均值

	fmt.Println("原始数组：", arr)
	max, index := search.GetMaxInts(arr)
	avg := search.GetAvg(arr)

	mySort.BubbleSort(&arr, "desc")
	fmt.Printf("倒序打印数组：%v\n", arr)

	fmt.Printf("最大值是%v, 下标为%v, 平均值为%v\n", max, index, avg)

	mySort.BubbleSort(&arr, "asc")
	search.BinarySearch(arr, 0, len(arr), 55)
}

//Q2 已知有个升序数组， 要求插入一个元素， 最后打印该元素，顺序依然是升序
func q2() {
	var arr = []int{3, 5, 8, 9, 15, 26}
	search.Insert(&arr, -5)
	search.Insert(&arr, 12)
	fmt.Println(arr)
}

//Q3 定义一个3行4列的二维数组， 逐个从键盘输入值， 编写程序将四周的数据清0
func q3() {
	const r, c = 3, 4
	var arr [3][4]int
	for i, rv := range arr {
		for j, _ := range rv {
			fmt.Printf("Input arr[%v][%v]'s val:", i, j)
			fmt.Scanln(&arr[i][j])
		}
	}
	arr = q3ZerofyBorder(arr)

	for _, vr := range arr {
		for _, vc := range vr {
			fmt.Printf("%v\t", vc)
		}
		fmt.Println()
	}
}

func q3ZerofyBorder(arr [3][4]int) [3][4]int {
	const r = 3 //3
	const c = 4 //4
	for i := 0; i < c; i++ {
		arr[0][i] = 0
		arr[r-1][i] = 0
	}
	for j := 0; j < r; j++ {
		arr[j][0] = 0
		arr[j][c-1] = 0
	}
	return arr
}

//Q4 定义一个4行4列的二维数组， 逐个从键盘输入值，然后将第一行和第四行的数据进行交换
func q4() {
	var arr [4][4]int

	for i, vr := range arr {
		for j, _ := range vr {
			fmt.Printf("Input arr[%v][%v]'s val:", i, j)
			fmt.Scanln(&arr[i][j])
		}
	}
	fmt.Println(arr)
	for j := 0; j < len(arr); j++ {
		arr[0][j], arr[len(arr)-1][j] = arr[len(arr)-1][j], arr[0][j]
	}

	for _, vr := range arr {
		for _, vc := range vr {
			fmt.Printf("%v\t", vc)
		}
		fmt.Println()
	}
}

//Q5 试保存1 3 5 7 9 五个奇数到数组，并倒序打印
func q5() {
	arr := search.GetOdd(10)
	mySort.BubbleSort(&arr, "desc")
	fmt.Println(arr)
}

//Q6 试写出实现查找的核心代码， 比如已知数组 mySort[10]string里面保存了10个元素， 现要查找
// “AA"在其中是否存在， 打印提示， 如果有多个"AA", 也要找到对应下标
func q6() {
	arr := []string{"AA", "AA", "Aa", "CDA", "Ca", "aB", "aBc"}
	sort.Strings(arr)
	search.BinarySearchString(arr, 0, len(arr), "AA")
}

//Q7 编写一个函数，可以接收一个数组，该数组有5个数，请找出最大的数和最小的数和对应的数组下标
func q7(arr []int) {
	max, maxIndex := search.GetMaxInts(arr)
	min, minIndex := search.GetMinInts(arr)

	fmt.Println(max, maxIndex, min, minIndex)
}

//Q8 定义一个数组， 并给出8个整数, 求该数组中大于平均值的数的个数，和小于平均值数的个数
func q8(arr []int) {
	avg := search.GetAvg(arr)
	greater := 0
	smaller := 0

	for _, v := range arr {
		if float64(v) > avg {
			greater++
		}
		if float64(v) < avg {
			smaller++
		}
	}
	fmt.Println("count > :", greater)
	fmt.Println("count < :", smaller)
}

//9 （有难度）
//跳水比赛，8个评委打分，运动员成绩是8个成绩去掉一个最高分，去掉一个最低分，剩下6个数算平均值
// 1。 打印最高分和最低分的评委
// 2。 找出最佳评委和最差评委， 最佳评委是打分和最后分数最接近的评委，最差评委反之。
func q9() {
	var scores [10]float64 = [10]float64 {1,2,3,4,5,6,7,8,9,0}
	high, hIdx := search.GetMaxFloats(scores)
	low, lIdx := search.GetMinFloats(scores)
	for i, _ := range hIdx {
		fmt.Printf("打分最高%v分的评委是%v号评委\n", high, hIdx[i] + 1)
	}
	for i, _ := range lIdx {
		fmt.Printf("打分最低%v分的评委是%v号评委\n", low, lIdx[i] + 1)
	}
	//去掉最高分，最低分
	scoresTmp := scores
	scoresTmp[hIdx[0]] = 0
	scoresTmp[lIdx[0]] = 0

	//算出最终平均分
	res := search.GetAvgFloats(scoresTmp)
 var ratio [10] float64
	for i, v := range scores {
		ratio[i] = math.Abs (1 - v / res) // ratio越小 越接近最终评分
	}
	_, bestIdx := search.GetMinFloats(ratio)
	_, worstIdx := search.GetMaxFloats(ratio)
	fmt.Println(ratio)
	for i, _ := range bestIdx {
		fmt.Printf("最佳评委是%v号评委\n", bestIdx[i] + 1)
	}
	for i, _ := range worstIdx {
		fmt.Printf("最差评委是%v号评委\n", worstIdx[i] + 1)
	}
}
func main() {

	//var arr = []int{3, 5, 52, 9, -7, 26}
	//q8(arr)
	q9()
}
