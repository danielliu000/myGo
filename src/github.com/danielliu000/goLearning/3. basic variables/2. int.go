package main

import (
	"fmt"
	"unsafe"
)

func main() {

	//int8 int16 int32 int64
	//uint8 uint16 uint32 uint64
	//int 有符号 uint 无符号（含0正数）
	//byte 相当于uint8 范围0-255 但使用的是unicode编码
	//uintptr 没有指定大小 可以足够保存指针

	var a byte = 'a'
	var age byte = 33 // 尽量声明符合数据范围的类型
	//-------------------------------------------------------
	var maxNum1 int8 = 127
	var maxNum2 int8 = -128
	fmt.Println(maxNum1 + 1) // 128 数据溢出 变成 -128
	fmt.Println(maxNum2 - 1) // -129 数据溢出 变成 127
	//--------------------------------------------------------
	n1 := 100
	//fmt.Printf 格式化输出
	fmt.Printf("n1的数据类型是： %T, " +"占用的字节为: %d\n", n1, unsafe.Sizeof(n1))
	fmt.Printf("a的类型: %T, 占用字节: %d", a, unsafe.Sizeof(a))
	fmt.Println(age)
}
