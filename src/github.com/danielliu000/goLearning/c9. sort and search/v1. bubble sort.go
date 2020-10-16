package main

import "fmt"

//func bubbleSort1(mySort []int) []int { // small => large
//
//	for j := 0; j < len(mySort); j++ {
//		for i := len(mySort) - testing; i > 0; i-- {
//			if mySort[i] < mySort[i-testing] {
//				mySort[i], mySort[i-testing] = mySort[i-testing], mySort[i]
//			}
//		}
//	}
//	return mySort
//}
func bubbleSort2(arr *[]int)  { // small => large

		arrLen := len(*arr)
		for i := 0; i < arrLen; i++ { //每次大循环最大值已经移到最右,不用再比了
			for j := 0; j < arrLen - 1 - i; j ++ {
					if (*arr)[j] > (*arr)[j + 1] {
						(*arr)[j], (*arr)[j + 1] = (*arr)[j + 1], (*arr)[j]
					}
			}
		}
}
func main() {
	arr := []int {33,12,2,-13,3,456,23,10}
	//fmt.Println(bubbleSort1(mySort))

	bubbleSort2(&arr)
	fmt.Println(arr)

}
