package main

import "fmt"

type Cat struct {
	Name string
	Age int
}
func main() {
	allChan := make(chan interface{}, 3)

	allChan <- 10
	allChan <- "tom"
	cat := Cat{"xiaohua", 4}
	allChan <- cat

	//想要获取第3个元素， 推出前2个
	<-allChan
	<-allChan

	newCat := <- allChan  //newCat type ?
	fmt.Printf("newCat's type: %T, newCat's value: %v\n", newCat, newCat)

	//此时newCat运行时看似是一个结构体，但编译时是一个接口，newCat.Name => 取不到Name, 因为没有实现接口方法
	//fmt.Printf("newCat's Name: %v", newCat.Name) //报错 type interface {} is interface with no methods
	//解决上述问题，需要添加类型断言
	c := newCat.(Cat)
	fmt.Printf("newCat's Name: %v", c.Name) //断言后可以取出


}
