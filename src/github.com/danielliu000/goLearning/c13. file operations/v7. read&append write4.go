package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

func main() {
	const PATH string = "E:/Daniel Files/IT_Learning/mygo/" +
		"src/github.com/danielliu000/goLearning/c13. file operations/test.txt"


	file, err := os.OpenFile(PATH, os.O_RDWR, 0666)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer file.Close()
	rw := bufio.NewReadWriter(bufio.NewReader(file), bufio.NewWriter(file))
	for {
		str, err := rw.ReadString('\n')
		if err == io.EOF  {
			break
		}
		fmt.Println(str)
	}

	_, err = rw.WriteString("FEDCBA\r\n")
	if err != nil {
		fmt.Println(err)
		return
	}
	err = rw.Flush()
	if err != nil {
		fmt.Println(err)
		return
	}
}
