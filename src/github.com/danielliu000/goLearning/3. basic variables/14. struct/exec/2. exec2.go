package main

import (
	"fmt"
	"math"
)

//1. 声明一个结构体circle 字段为radius
//2. 声明一个方法area和circle绑定，可以返回面积
type Circle struct {
	Radius float64
}
func (circle *Circle) area () float64 {
	return math.Pi * math.Pow(circle.Radius ,2)
}
func main() {
 c := Circle{}
	c.Radius = 10
	fmt.Println(c.area())
}
