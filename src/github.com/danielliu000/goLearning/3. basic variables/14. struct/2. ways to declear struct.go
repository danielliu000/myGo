package main

import "fmt"

func main() {

	//声明结构体方式

	type Dog struct {
		Color string
		Age int
		Name string
	}
	//方式一
	var d1 Dog
	d1.Color = "white"
	d1.Age = 2
	d1.Name = "wangwang"
	//方式二
	d2 := Dog {"black", 23, "xiaowang"}
	fmt.Println(d2)
	//方式三
	var d3 *Dog = new(Dog) // d3是指向Dog的指针
	// d3 := new(Dog)
	//标准写法如下：
	(*d3).Name = "dawang"
	(*d3).Age = 3
	(*d3).Color = "brown"
	fmt.Println(*d3)

	//好像太啰嗦 等价写法如下: 指针.属性 = value
	//go底层会自动处理 指针.属性 的写法
	d3.Name = "dadawang"
	d3.Age = 4
	d3.Color = "Pink"
	fmt.Println(*d3)

	//方式四
	d4 := &Dog {Color: "aaa", Age: 3}  //还是一个指针 可以赋初值
	d4.Name = "xxx"
	d4.Age = 5
	fmt.Println(*d4)

}
