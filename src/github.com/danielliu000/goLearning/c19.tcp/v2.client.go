package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
	"strings"
)

func main() {
	conn, err := net.Dial("tcp", "0.0.0.0:8888")
	if err != nil {
		fmt.Println("connection failed", err)
		return
	}
	fmt.Println("连接成功")
	//从终端读取一行用户输入，并准备发送给服务器
	reader := bufio.NewReader(os.Stdin)
	for {
		str, err := reader.ReadString('\n')

		if err != nil {
			fmt.Println("readString err", err)
		}
		str = strings.Trim(str, "\n\r")
		if str == "exit" {
			fmt.Println("客户端退出")
			break
		}
		//将str发送给服务器
		n, err := conn.Write([]byte(str + "\n"))
		if err != nil {
			fmt.Println("conn.Write", err)
		}
		fmt.Printf("客户端发送了%d字节的数据\n", n)
	}

}
