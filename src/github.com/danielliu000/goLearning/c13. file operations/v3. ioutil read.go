package main

import (
	"fmt"
	"io/ioutil"
)

//一次性读取 适合文件不大的情况
//不需要显式打开关闭， 文件打开关闭封装到ReadFile中了
func main() {
	file := "E:/test.txt"
	content, err := ioutil.ReadFile(file) //content 为 []byte
	if err != nil {
		fmt.Println("Read file error: ", err)
	}

	fmt.Printf("%v",string(content))

}
