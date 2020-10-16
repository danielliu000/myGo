package main

import "fmt"

func binarySearch1(arr []int) {
	leftIndex := 0
	rightIndex := len(arr) - 1
	midIndex := 0
	num := 0
	fmt.Println("Search for: ")
	_, _ = fmt.Scanln(&num)

Loop:
	for {
		if leftIndex > rightIndex {
			fmt.Println("找不到")
			break Loop
		}
		midIndex = (leftIndex + rightIndex) / 2
		switch {
		case num == arr[midIndex]:
			fmt.Printf("找到%v, 下标是%v\n", num, midIndex)
			break Loop

		case num < arr[midIndex]:
			rightIndex = midIndex - 1

		case num > arr[midIndex]:
			leftIndex = midIndex + 1
		}
	}

} //不使用递归

func binarySearch2(arr *[]int, left int, right int, searchVal int) {
	if left > right {
		fmt.Println("没找到")
		return
	}
	mid := (left + right) / 2

	if (*arr)[mid] > searchVal {
		binarySearch2(arr, left, mid-1, searchVal)
	} else if (*arr)[mid] < searchVal {
		binarySearch2(arr, mid+1, right, searchVal)
	} else {
		fmt.Printf("找到了, 下标为%v \n", mid)

	}
} //使用递归
func main() {
	//二分查找 Binary Search 要求: 有序序列
	arr := []int{-9, 3, 19, 20, 54, 54, 54, 111, 121, 141, 211}
	binarySearch1(arr)
	//binarySearch2(&mySort, 0, len(mySort), -9)

}
