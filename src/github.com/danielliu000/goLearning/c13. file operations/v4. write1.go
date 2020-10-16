package main

import (
	"bufio"
	"fmt"
	"os"
)
type myConst struct {

}
const PATH string = "E:/Daniel Files/IT_Learning/mygo/" +
"src/github.com/danielliu000/goLearning/c13. file operations/test.txt"

//testing. 写入5句 Hello, Daniel 到一个文件，没有改文件则创建
func main() {

	file, err := os.OpenFile(PATH, os.O_CREATE|os.O_WRONLY, 0666)
	defer file.Close()
	if err != nil {
		fmt.Println("open file error: ", err)
		return
	}
	str := "Hello, Daniel\r\n" //不同编辑器换行符不同 \r \n并用保证换行
	writer := bufio.NewWriter(file)

	for i := 0; i < 5; i++ {
		_, err = writer.WriteString(str)
		if err != nil {
			fmt.Println("write file error: ", err)
		}

	}
	err = writer.Flush()
	if err != nil {
		fmt.Println("write file error: ", err)
	}
}
