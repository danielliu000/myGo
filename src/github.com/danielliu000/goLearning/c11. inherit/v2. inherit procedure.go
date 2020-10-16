package main

import "fmt"

type A struct {
	Name string
}

func (a *A) showA() {
	fmt.Println("学生姓名: ", a.Name)
}

type B struct {
	A
	Name string
}

func (b *B) showB() {
	fmt.Println("学生姓名: ", b.Name)
}
type C struct {
	A
	B
}

type D struct {
	DB A   //有名结构体
	int
}


func main() {
	b := B{
		A:    A{"Daniel"},
		Name: "Winnie",
	}

	b.showB() //优先查找自身结构体的属性，如果没有再向上查找
	b.showA() //b没有绑定showA方法，因此会找到结构体A中的方法，A查找自身属性，输出..

	c := C{
		A: A{},
		B: B{},
	}


	//c.Name = "leo" //错误 A,B 中都有Name属性，赋值时必须指定，否则编译错误
	c.A.Name = "Leo" //正确


	d := D{}
	//d.Name = "John"  //错误 有名结构体赋值时编译器只看自身属性，没有会报错，所以必须指定属性名
	d.DB.Name = "John" //正确
	d.int = 40  // 基本类型也可以嵌入结构体
}
