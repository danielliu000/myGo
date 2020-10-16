package main

import (
	"bufio"
	"fmt"
	"os"
)
//对文件追加内容
func main() {
	const PATH string = "E:/Daniel Files/IT_Learning/mygo/" +
		"src/github.com/danielliu000/goLearning/c13. file operations/test.txt"
	file, err := os.OpenFile(PATH, os.O_APPEND, 0666)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer file.Close()
	writer := bufio.NewWriter(file)
	_, err = writer.WriteString("ABCDEFG\r\n")
	if err != nil {
		fmt.Println(err)
		return
	}
	err = writer.Flush()
	if err != nil {
		fmt.Println(err)
		return
	}
}
