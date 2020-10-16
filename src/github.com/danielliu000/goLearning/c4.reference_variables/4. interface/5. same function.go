package main

import "fmt"
//接口AB BC 有重复的test02() AC继承了AB, BC 但不会报错
type AB interface {
	test01()
	test02()
}

type BC interface {
	test02()
	test03()
}

type AC interface {
	AB
	BC
}

type W struct {

}

func (w *W) test01() {fmt.Println("test01")}
func (w *W) test02() {fmt.Println("test02")}
func (w *W) test03() {fmt.Println("test03")}
func main() {
	w := W{}
	var a AC = &w

	a.test01()
	a.test02()
	a.test03()



}
