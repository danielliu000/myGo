package main

import "fmt"

func main() {
	var arr2 [3][5]int
	totalSum := 0.0
	var classSumAvg [3][2]float64

	for i, v := range arr2 {
		sum := 0
		fmt.Printf("class %v\n", i+1)
		for j, v2 := range v {
			fmt.Printf("input score for student %v: ", j+1)
			_, _ = fmt.Scanln(&v2)
			arr2[i][j] = v2
			sum += v2
		}
		//存储每班总分，平均分
		classSumAvg[i][0] = float64(sum)
		classSumAvg[i][1] = float64(sum) / float64(len(arr2[1]))
	}

	for i, v := range classSumAvg {
		fmt.Printf("class %v's total score is %v, average is %.2f\n",
			i+1, v[0], v[1])
		totalSum += v[0]
	}
	fmt.Printf("Grade total score is %v, average is %.2f",
		totalSum, totalSum/float64(len(arr2)*len(arr2[1])))

}
