package main

import (
	"chatRoom/server/main/processor"
	"chatRoom/server/main/redis"
	"chatRoom/server/model"
	"fmt"
	"net"
)

func handleProcess(conn net.Conn) {
	//延时关闭conn
	//defer conn.Close()
	//循环读取客户端发送的信息
	process := &processor.Processor{
		Conn:conn,
	}
	err := process.ProcessHandler()
	if err != nil {
		fmt.Println("Server Client 协程错误:  ", err)
		return
	}


}
//对userDao初始化
func initUserDao() {
	model.MyUserDao = model.NewUserDao(redis.Rdb)
	fmt.Println(redis.Rdb)
	fmt.Println(model.MyUserDao)
}
func main() {
	//初始化Redis连接

	redis.InitRedisClient()
	initUserDao()
	fmt.Println("服务器在8889端口监听...")
	listener, err := net.Listen("tcp", "0.0.0.0:8889")
	if err != nil {
		fmt.Printf("net.Listen Error: %v\n", err)
		return
	}
	defer listener.Close()

	for {
		fmt.Println("等待客户端连接....")
		conn, err := listener.Accept()
		if err != nil {
			fmt.Printf("listener.Accept Error: %v\n", err)
		}
		//连接成功后，启动协程和客户端保持通信
		go handleProcess(conn)
	}

}
