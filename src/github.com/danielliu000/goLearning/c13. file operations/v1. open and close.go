package main

import (
	"fmt"
	"os"
)

func main() {
	//open a file
	file, err := os.Open("E:/test.txt")
	if err != nil {
		fmt.Println("open file error: ", err)
	}
	fmt.Printf("file=%v",file) //文件file是一个指针
	defer file.Close()

}
