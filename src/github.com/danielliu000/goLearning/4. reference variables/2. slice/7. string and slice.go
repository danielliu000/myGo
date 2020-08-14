package main

import "fmt"

func main() {
	//1. string底层是一个byte数组,因此string也可进行切片梳理
	//2. string 不可变
	//3. 改变string 先要将string转化成[]byte或[]rune 再重写string

	str := "Hello World@Daniel"
	strSlice := str[:11]
	fmt.Println(strSlice)
	//str[0] = "z"  不可修改
	strByte := []byte(str)  // []byte 无法改写中文
	strByte[0] = 'B'
	strB := string(strByte)

	strRune := []rune(str) // []rune 可以改写中文
	strRune[1] = '刘'
	strR := string(strRune)


	fmt.Println(str)
	fmt.Printf("strByte=%c, strB=%v\n", strByte, strB)
	fmt.Printf("strByte=%c, strR=%v", strRune, strR)




}
