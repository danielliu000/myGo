package main

import "fmt"

//testing. 编写结构体(MethodUtils) 编程一个方法，方法不需要参数， 在方法中打印一个10 * 8的矩形
// 在main方法中调用。
//2. 编写一个方法， 提供m和n二个参数 方法中打印一个 m*n的矩形
//3. 编写一个方法计算该矩形的面积（可以接收长len， 宽width)， 将其作为方法返回值。在main方法
// 中调用该方法，接收返回的面积值并打印。
type MethodUtils struct {

}

//q1
func (mu *MethodUtils) q1() {
	for i := 1; i <= 10; i++ {
		for j := 1; j <= 8; j++ {
			fmt.Print("* ")
		}
		fmt.Println()
	}
}
//q2
func (mu *MethodUtils) q2(l int, w int) {
	for i := 1; i <= l; i++ {
		for j := 1; j <= w; j++ {
			fmt.Print("* ")
		}
		fmt.Println()
	}
}

//q3
func (mu *MethodUtils) q3(len float64, width float64)  float64 {
	return len * width
}
func main() {
	m := MethodUtils{}
	//m.q1()
	m.q2(2, 6)
	fmt.Println("面积为: ", m.q3(3, 4))

}
