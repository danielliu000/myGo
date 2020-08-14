package search

import (
	"utils/mySort"
)

func Insert(arr *[]int, insertVal int) {
	//mySort.BubbleSort(arr, "asc")
	slice := make([]int, 0, len(*arr) + 1)
	slice = append(slice, *arr...)
	slice = append(slice, insertVal)
	mySort.BubbleSort(&slice, "asc")
	*arr = slice
}
