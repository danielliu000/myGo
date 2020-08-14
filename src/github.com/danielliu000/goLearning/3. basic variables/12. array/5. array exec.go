package main

import (
	"fmt"
	"math/rand"
	"time"
)

//1. 数组存放26个大写字母
func saveLetters () {
	arr := [26]byte{}
	for i := 0; i < len(arr); i++ {
		arr[i] = 'A' + byte(i)
	}
	fmt.Printf("%c", arr)
}

//2. 求数组中最大值， 并打印下标
func getMax(arr [10]int) int {
	max := arr[0]
	for i := 0; i < len(arr); i++ {
		if arr[i] > max {
			max = arr[i]
		}
	}
	return max
}

//3. 求数组的和，和 平均值

func getAverage(arr [10]int) (int, float64) {
	sum := 0
	for _, v := range arr {
		sum += v
	}
	return sum, float64(sum) / float64(len(arr))
}

//4. 随机生成5个数，反转打印

func getRandomNum() {
	arr := [10]int{}
	seed := time.Now().UnixNano()
	rand.Seed(seed)
	for i := len(arr) - 1; i >= 0; i-- {
		randNum := rand.Intn(100)
		fmt.Printf("%v ",randNum)
		arr[i] = randNum
	}
	fmt.Println()
	for _, v := range arr {
		fmt.Printf("%v ", v)
	}
}
func main() {
	//saveLetters()
	arr := [10]int {21,245,23,54,643,34,76,233,23,2}
	fmt.Println(getMax(arr))
	fmt.Println(getAverage(arr))
	getRandomNum()
}
