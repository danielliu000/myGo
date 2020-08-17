package main

import "fmt"

//空接口没有任何方法， 所以所有类型都实现了空接口，
//即我们可以把任何一个变量赋给空接口
type T interface {
}

type A struct {

}
func main() {
	a := A{}
	var t1 T = a
	fmt.Println(t1)

	var t2 interface{} = a
	num1 := 8.8

	t2 = num1
	t1 = num1

	fmt.Println(t1, t2)


}
