package main

import "fmt"

type Point struct {
	x int
	y int
}
func main() {
	var a interface{}
	var point Point = Point{1,2}
	a = point
	var b Point
	//b = a // 不可以直接赋值 需要断言
 //类型断言，
 //判断a是否指向Point类型的变量，
 //如果是则转换成Point并赋值为b，否则报错。

	if b, ok := a.(Point); ok {
		fmt.Println(b)
	} else {
		fmt.Println("convert failed")
	}

}
