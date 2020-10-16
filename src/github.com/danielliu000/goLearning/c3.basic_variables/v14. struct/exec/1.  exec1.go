package main

import "fmt"

type Person struct {
	Name string  `json:"name"`
	Age byte			`json:"age"`
	Gender string `json:"gender"`
}
//testing. 给结构体Person 增加一个speak方法，输出 "xxx是一个好人"

func (person Person) speak() {
	person.Name = "Daniel"
	fmt.Println(person.Name, "is a good man")
}
//2. 给Person 添加一个计算方法 计算从1- 1000 的和
func (person Person) getSum( num int) int {
	sum := 0
	for i := 1; i <= num; i++ {
		sum += i
	}
	return sum
}
func main() {
	p := Person{}
	p.speak()
	fmt.Println(p.getSum(1000))
}
