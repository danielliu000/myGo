package search

func GetMaxInts(arr []int) (int, []int) {
	max := arr[0]
	var index []int

	//得到最大值
	for _,v := range arr {
		if v >= max {
			max = v
		}
	}
	//得到最大值所有下标
	for i, v := range arr {
			if v == max {
				index = append(index, i)
			}
		}
	return max, index
}
func GetMaxFloats(arr [10]float64) (float64, []int) {
	max := arr[0]
	var index []int

	//得到最大值
	for _,v := range arr {
		if v >= max {
			max = v
		}
	}
	//得到最大值所有下标
	for i, v := range arr {
			if v == max {
				index = append(index, i)
			}
		}

	return max, index
}