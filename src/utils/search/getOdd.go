package search

func GetOdd(num int) []int {
	arr := make([]int, 0)
	for i := 0; i < num; i++ {
		if i % 2 == 1 {
			arr = append(arr, i)
		}
	}
	return arr
}