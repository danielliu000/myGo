package main

import (
	"bufio"
	"fmt"
	"os"
)

//2. 写入10句Hello world！ 替换原文件内容
func main() {
	const PATH string = "E:/Daniel Files/IT_Learning/mygo/" +
		"src/github.com/danielliu000/goLearning/c13. file operations/test.txt"
	file, err := os.OpenFile(PATH, os.O_WRONLY | os.O_TRUNC, 0666)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer file.Close()
	writer := bufio.NewWriter(file)
	str := "Hello World!\r\n"
	for i := 0; i < 10; i++ {
		_, err = writer.WriteString(str)
		if err != nil {
			fmt.Println(err)
			return
		}
	}

	err = writer.Flush()
	if err != nil {
		fmt.Println(err)
		return
	}
}
