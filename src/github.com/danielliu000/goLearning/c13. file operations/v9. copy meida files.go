package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

//io copy media


func copyFile(dstFilePath, srcFilePath string) (written int64, err error ){

	srcFile, e := os.Open(srcFilePath)
	if e != nil {
		fmt.Println(e)
	}
	dstFile, e := os.OpenFile(dstFilePath, os.O_CREATE | os.O_WRONLY, 0666)
	if e != nil {
		fmt.Println(e)
	}
	defer srcFile.Close()
	defer dstFile.Close()

	rw := bufio.NewReadWriter(bufio.NewReader(srcFile), bufio.NewWriter(dstFile))

	return io.Copy(rw.Writer, rw.Reader)
}
func main() {
	srcPath := "./2.jpg"
	dstPath := "E:/Daniel Files/IT_Learning/mygo/" +
		"src/github.com/danielliu000/goLearning/c13. file operations/2copy.jpg"
	_, err := copyFile(dstPath, srcPath)
	if err != nil {
		fmt.Println(err)
	}
}
