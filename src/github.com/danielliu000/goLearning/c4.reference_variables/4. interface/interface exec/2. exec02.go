package main

import (
	"fmt"
	"math/rand"
	"sort"
)

//练习： 与练习1类似， 实现Student切片对Score从大到小排序sort.Sort(data interface)
type Student struct {
	Name string
	Age  int
	Score float64
}

type StuSlice []Student

func (stu StuSlice) Len() int {
	return len(stu)
}
func (stu StuSlice) Less(i, j int) bool {
	return stu[i].Score > stu[j].Score
}
func (stu StuSlice) Swap(i, j int) {
	stu[i], stu[j] = stu[j], stu[i]
}

func main() {
	var students StuSlice
	for i := 0; i < 10; i++ {
		student := Student{
			Name:  fmt.Sprintf("student-%02d", i + 1),
			Age:   rand.Intn(30),
			Score: float64(rand.Intn(100)) + rand.Float64(),
		}
		students = append(students, student)
	}

	fmt.Println("---------排序前------------")
	for _, v := range students {
		fmt.Println(v)
	}
	sort.Sort(students)
	fmt.Println("---------排序后------------")
	for _, v := range students {
		fmt.Printf("%v, Age: %v, Score: %.2f\n", v.Name, v.Age, v.Score)
	}

}
