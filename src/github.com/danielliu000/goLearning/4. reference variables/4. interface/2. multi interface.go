package main

import "fmt"

//一个自定义类型可以实现多个接口
type AInterface interface {
	Hello()
}
type BInterface interface {
	Say()
}

type Student struct {

}

func( stu *Student) Hello() {
	fmt.Println("Student, hello()")
}
func( stu *Student) Say() {
	fmt.Println("Student, Say()")
}

func main() {

	s := Student{}
	//student实现了AInterface, BInterface
	s.Hello()
	s.Say()
}
