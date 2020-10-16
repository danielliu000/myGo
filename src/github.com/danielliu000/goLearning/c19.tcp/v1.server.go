package main

import (
	"fmt"
	"net"
)
func process(conn net.Conn) {
	//循环接收客户端发送的数据
	defer conn.Close() //关闭conn
	for {
		//创建一个切片
		buf := make([]byte, 1024)
		//1.等待客户端通过conn发送信息
		//2.如果客户端没有write[发送]，那么协程就阻塞在这里
		//fmt.Printf("服务器在等待客户端%s发送信息\n", conn.RemoteAddr().String())
		n, err := conn.Read(buf) //从conn读取
		if err != nil {
			fmt.Println("客户端退出,", err)
			return
		}
		//3. 显示客户端发送的内容到服务器终端
		fmt.Print(string(buf[:n]))
	}

}
func main() {
	listener, err := net.Listen("tcp", "0.0.0.0:8888")
	if err != nil {
		fmt.Println("connection failed", err)
		return
	}

	defer listener.Close()
	for {
		//fmt.Println("等待客户端连接...")
		conn, err := listener.Accept()
		if err != nil {
			fmt.Println(err)
		} else {
			//fmt.Println(conn)
			//fmt.Println("客户端地址：", conn.RemoteAddr().String())
			//开启协程
			go process(conn)
		}
	}
}
