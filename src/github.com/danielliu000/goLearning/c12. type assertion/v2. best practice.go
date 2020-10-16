package main

import "fmt"

//对前面Usb接口案例做出改进
//给Phone结构体增加一个特有的方法call(),
//当Usb，接口接收的是Phone变量时，还需要调用call方法
type Usb interface {
	Connect()
	Disconnect()
}
type Phone struct { Brand string}
func (p *Phone) Connect() {
	fmt.Println(p.Brand, " is connected")
}
func (p *Phone) Disconnect() {
	fmt.Println(p.Brand, " is disconnected")
}

//call 方法时Phone结构体特有的
func (p *Phone) Call() {
	fmt.Println(p.Brand, " starts calling...")
}
type Camera struct {Brand string}
func (c *Camera) Connect() {
	fmt.Println(c.Brand, " is connected")
}
func (c *Camera) Disconnect() {
	fmt.Println(c.Brand, " is disconnected")
}
type Computer struct {}
func (c *Computer) Working(usb Usb) {
	usb.Connect()
	//这里进行类型断言,如果是phone类型则调用call方法
	if phone, ok := usb.(*Phone); ok {
		phone.Call()
	}
	usb.Disconnect()


}
func main() {

	var usbArr [3]Usb
	//目标： 遍历usbArr 如果时Phone 结构体 则需要自动调用其特有的call()方法

	usbArr[0] = &Phone{"Huawei"}
	usbArr[1] = &Phone{"Vivo"}
	usbArr[2] = &Camera{"Canon"}

	computer := &Computer{}

	for _, v := range usbArr {
		computer.Working(v)
		fmt.Println()
	}







}
