package main

import (
	"fmt"
	"io/ioutil"
)

//将一个文件内容复制到另一个文件
//使用ioutil
func main() {
	const PATH1 string = "E:/Daniel Files/IT_Learning/mygo/" +
		"src/github.com/danielliu000/goLearning/c13. file operations/test.txt"
	const PATH2 string = "E:/Daniel Files/IT_Learning/mygo/" +
		"src/github.com/danielliu000/goLearning/c13. file operations/test2.txt"
	content, err := ioutil.ReadFile(PATH2)
	if err != nil {
		fmt.Println(err)
		return
	}
	err = ioutil.WriteFile(PATH1,content,0666)
	if err != nil {
		fmt.Println(err)
		return
	}
}
