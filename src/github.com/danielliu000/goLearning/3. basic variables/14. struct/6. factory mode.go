package main

import (
	"fmt"
	"model/person"
)
//使用工长函数解决小写结构体无法引入问题

func main() {
	//student := utils.student { // student 结构体是小写 无法从外部引入怎么办
	//	Name: "Daniel",
	//	Age: 33,
	//	Gender: "male",
	//}
	//调用工厂模式函数

	//此时student是一个结构体指针
	student := person.NewStudent("Winnie", 34, "female")
	fmt.Println(*student)
	fmt.Println(student.Name) //取值时，可以不用加*
	fmt.Println(student.Age) //取值时，可以不用加*
	fmt.Println(student.Gender) //取值时，可以不用加*
}
