package main

import "fmt"

func main() {
	var a byte = 'a'  //单引号-字符 双引号-字符串
	var b byte = '0'
	//var c byte = '刘'  //overflow 溢出
	var d int16 = '刘'  // int16 长度可以不会溢出

	fmt.Println("a=",a, "b=",b) //输出码值
	fmt.Printf("a=%c b=%c\n", a, b) //格式化输出%c 对应字符
	fmt.Printf("d=%c\n", d) //格式化输出%c 对应字符

	for i := 0; i < 26; i++ {
		fmt.Printf("Letter: %c\n", i + 'a' ) //输出26个字母
	}
}


