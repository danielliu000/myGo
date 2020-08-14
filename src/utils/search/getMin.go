package search


func GetMinInts(arr []int) (int, []int) {
	min := arr[0]
	var index []int

	//得到最小值
	for _,v := range arr {
		if v < min {
			min = v
		}
	}
	//得到最小值所有下标
	for i, v := range arr {
		if v == min {
			index = append(index, i)
		}
	}

	return min, index
}
func GetMinFloats(arr [10]float64) (float64, []int) {
	min := arr[0]
	var index []int

	//得到最小值
	for _,v := range arr {
		if v < min {
			min = v
		}
	}
	//得到最小值所有下标
	for i, v := range arr {
		if v == min {
			index = append(index, i)
		}
	}

	return min, index
}