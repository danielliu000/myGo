package search

import (
	"fmt"
)

func BinarySearchString(arr []string, firstIndex int, lastIndex int, searchVal string) {
	midIndex := (firstIndex + lastIndex) / 2

	if firstIndex > lastIndex {
		fmt.Println(searchVal, "没有找到")
		return
	}
	if searchVal == arr[midIndex] {
		fmt.Printf("找到了%v, 下标为%v\n", searchVal, midIndex)
		for i := midIndex + 1; i < lastIndex; i++ {
			if searchVal != arr[i] {
				break
			}else {
				fmt.Printf("找到了%v, 下标为%v\n", searchVal, i)
			}
		}
		for i := midIndex - 1; i >= 0; i--{
			if searchVal != arr[i] {
				break
			}else {
				fmt.Printf("找到了%v, 下标为%v\n", searchVal, i)
			}
		}

	} else if searchVal > arr[midIndex] {
		BinarySearchString(arr, midIndex + 1, lastIndex, searchVal)
	} else if searchVal < arr[midIndex] {
		BinarySearchString(arr, firstIndex, midIndex - 1, searchVal)
	}
}
