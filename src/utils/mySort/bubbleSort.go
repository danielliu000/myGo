package mySort

func BubbleSort(arr *[]int,  order string) { //order : desc / asc
	arrLen := len(*arr)
	for i := 0; i < arrLen; i++ {
		for j := 0; j < arrLen - 1- i; j++ {
			if order == "asc" && (*arr)[j] > (*arr)[j + 1] {
				(*arr)[j] , (*arr)[j + 1] = (*arr)[j + 1], (*arr)[j]
			} else if order == "desc" && (*arr)[j] < (*arr)[j + 1] {
				(*arr)[j] , (*arr)[j + 1] = (*arr)[j + 1], (*arr)[j]
			}
		}
	}

}
