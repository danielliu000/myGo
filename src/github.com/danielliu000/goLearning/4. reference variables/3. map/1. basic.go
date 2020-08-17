package main

import "fmt"

func main() {
	//方式一： 先声明，再make
	var a map[int]string
	//使用map前需要make, 来给map分配数据空间
	a = make(map[int]string, 10)

	a[1] = "Daniel"
	fmt.Println(a)

	//方式二： 声明同时make
	b := make(map[string]string)
	b["0"] = "0"
	fmt.Println(b)

	//方式三：声明同时赋值
	c := map[string]string {
		"0": "0",
		"1": "1",
	}
	fmt.Println(c)
}
