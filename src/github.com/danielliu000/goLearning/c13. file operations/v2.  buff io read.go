package main
//带缓冲区的读取
import (
	"bufio"
	"fmt"
	"io"
	"os"
)

func main() {
	//open a file
	file, err := os.Open("E:/test.txt")
	if err != nil {
		fmt.Println("open file error: ", err)
	}
	fmt.Printf("file=%v\n",file) //文件file是一个指针
	defer file.Close() //延迟关闭文件
	reader := bufio.NewReader(file)

	for {
		str, err := reader.ReadString('\n')// 读到一个换行就结束
		if err == io.EOF { //文件末尾
			break
		}
		fmt.Print(str)
	}

}
