package main

import "fmt"

type Usb interface {
	//接口默认是一个指针
	//接口中不能有方法体和变量
	//体现了程序设计的多态和高内聚低偶合的思想
	Start()
	Stop()
}

type Phone struct {}

func (p Phone) Start() {
	fmt.Println("phone starts working")
}

func (p Phone) Stop() {
	fmt.Println("phone stops working")
}

type Camera struct {}


func (c Camera) Start() {
	fmt.Println("Camera starts working")
}

func (c Camera) Stop() {
	fmt.Println("Camera stops working")
}
type Computer struct {

}


//Working 方法，接收一个Usb接口类型变量

//实现Usb接口，就是指实现了Usb 接口声明所有方法
func (computer *Computer) Working(usb Usb) {
	usb.Start()
	usb.Stop()
}
func main() {
	computer := Computer{}
	phone := Phone{}
	camera := Camera{}


	//调用接口
	computer.Working(phone)
	computer.Working(camera)
}
