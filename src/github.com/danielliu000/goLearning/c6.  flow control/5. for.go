package main

import "fmt"

func main() {
	for i:=0; i<10; i++ {
		fmt.Println(i)
	}
	//遍历字符串
	//方式一 传统方式
	str := "hello world，悉尼"
	strRune := []rune(str)
	strLen := len(strRune)
	for i := 0; i < strLen; i++ {
		//传统方式中文会出现乱码，原因此方法遍历是按字节遍历，而中文一个汉字是3个字节
		//解决方案 将str变成切片 []rune
		fmt.Printf("%c\n", strRune[i])
	}

	//方式二  for ... range
	for index, val := range str {
		//
		fmt.Printf("index: %d, val: %c\n", index, val)
	}

}
