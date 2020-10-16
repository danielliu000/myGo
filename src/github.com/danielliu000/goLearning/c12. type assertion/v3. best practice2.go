package main

import (
	"fmt"
)

//编写一个函数，循环判断传入参数类型
func typeJudge(items... interface{}) {
	for i, v := range items {
		switch v.(type) {
		case int,byte, int16, int32, int64:
			fmt.Printf("第%v个参数的类型是int,值是%v.\n", i, v)
		case bool:
			fmt.Printf("第%v个参数的类型是bool,值是%v.\n", i, v)
		case string:
			fmt.Printf("第%v个参数的类型是string,值是%v.\n", i, v)
		case float64,float32:
			fmt.Printf("第%v个参数的类型是float,值是%v.\n", i, v)
		case Student:
			fmt.Printf("第%v个参数的类型是Student,值是%v.\n", i, v)
		case *Student:
			fmt.Printf("第%v个参数的类型是*Student,值是%v.\n", i, v)
		default:
			fmt.Printf("第%v个参数的类型不确定,值是%v.\n", i, v)
		}
	}
}

type Student struct {

}
func main() {
	var n1 float32
	var n2 int64
	var n3 string
	var n4 bool
	var n5 uintptr
	var stu1 Student
	var stu2 *Student
	typeJudge(n1,n2,n3,n4,n5, stu1, stu2)
}
