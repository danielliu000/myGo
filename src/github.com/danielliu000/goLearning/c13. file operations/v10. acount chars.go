package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

//统计一个文件中英文、数字和其他字符的个数

type CharCount struct {
	ChCount int
	NumCount int
	SpaceCount int
	OtherCount int
}


func main() {
	const PATH string = "E:/Daniel Files/IT_Learning/mygo/" +
		"src/github.com/danielliu000/goLearning/c13. file operations/test.txt"
	file, err := os.Open(PATH)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer file.Close()
	var count CharCount
	reader := bufio.NewReader(file)

	for {
		str, err := reader.ReadString('\n')
		if err == io.EOF {
			break
		}
		for _, v := range str {
			switch  {
			case v >= 'a' && v <= 'z':
				count.ChCount++
			case v >= 'A' && v<= 'Z':
				count.ChCount++
			case v == ' ' || v == '\t':
				count.SpaceCount++
			case v >= '0' && v <= '9':
				count.NumCount++
			default:
				count.OtherCount++
			}
		}
	}

	fmt.Printf("字符个数=%v, 数字个数=%v,空格个数=%v, 其他字符=%v",
		count.ChCount, count.NumCount, count.SpaceCount, count.OtherCount)
}
