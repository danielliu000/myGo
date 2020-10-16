package process

import (
	"chatRoom/client/utils"
	"chatRoom/common/message"
	"encoding/json"
	"fmt"
	"net"
)

type UserProcess struct {

}

func(userProcess *UserProcess) Register(userId int,
	userPwd string, userName string)(err error) {
	//1. 连接到服务器
	conn, err := net.Dial("tcp", "localhost:8889")
	if err != nil {
		return err
	}

	defer conn.Close()

	registerMsg := message.RegisterMsg{
		User: message.User{
			UserId:   userId,
			UserPwd:  userPwd,
			UserName: userName,
		},
	}

	data, err := json.Marshal(registerMsg)
	if err != nil {
		return err
	}

	msg := message.Message{
		Type: message.RegisterMsgType,
		Data: string(data),
	}
	//序列化msg结构体
	data, err = json.Marshal(msg)
	if err != nil {
		fmt.Println("序列化msg失败: ",err)
		return
	}
	fmt.Println(string(data))

	tf := &utils.Transfer{
		Conn: conn,
		Buf:  [8096]byte{},
	}
	err = tf.WritePkg(data)
	//这里需要处理服务器端返回的消息
	if err != nil {
		fmt.Println("注册失败", err)
		return
	}
	msg, err = tf.ReadPkg()
	if err != nil {
		return
	}
	//将msg.Data部分反序列化成 RegisterResMsg
	var registerResMsg message.RegisterResMsg
	err = json.Unmarshal([]byte(msg.Data), &registerResMsg)

	if registerResMsg.StatusCode == 200 {
		fmt.Println("注册成功")

	} else if registerResMsg.StatusCode == 400 {
		fmt.Println(registerResMsg.Error)
	}
	return

}

func (userProcess *UserProcess) Login(userId int,
	userPwd string) (err error) {
	//fmt.Printf("userId=%d, userPwd=%s\n", userId, userPwd)
	//1. 连接到服务器
	conn, err := net.Dial("tcp", "localhost:8889")
	if err != nil {
		return err
	}

	defer conn.Close()
	//2. 通过conn给服务器发送消息
	loginMsg := message.LoginMsg{
		UserId:   userId,
		UserPwd:  userPwd,
		UserName: "",
	}
	//发送信息序列化
	data, err := json.Marshal(loginMsg)
	if err != nil {
		return err
	}

	msg := message.Message{
		Type: message.LoginMsgType,
		Data: string(data),
	}
	//序列化msg结构体
	data, err = json.Marshal(msg)
	if err != nil {
		fmt.Println("序列化msg失败: ",err)
		return
	}
	fmt.Println(string(data))

	//先发送长度，再发送msg
	tf := &utils.Transfer{
		Conn: conn,
		Buf:  [8096]byte{},
	}
	err = tf.WritePkg(data)
	//这里需要处理服务器端返回的消息

	msg, err = tf.ReadPkg()
	if err != nil {
		return
	}
	//将msg.Data部分反序列化成 LoginResMsg
	var loginResMsg message.LoginResMsg
	err = json.Unmarshal([]byte(msg.Data), &loginResMsg)

	if loginResMsg.StatusCode == 200 {
		//初始化curUser
		curUser.Conn = conn
		curUser.UserId = userId
		curUser.UserStatus = message.UserOnline

		//fmt.Println("登录成功")
		//显示在线用户列表
		fmt.Println("当前在线用户列表如下：")
		for _, v := range loginResMsg.UserIds {
			//不显示自己在线
			if v == userId {
				continue
			}
			fmt.Println("用户id:\t", v)
			//完成客户端onlineUsers初始化
			user := &message.User{
				UserId:   v,
				UserStatus: message.UserOnline,
			}
			onlineUsers[v] = user
		}
		fmt.Println()
		//启动一个协程,该协程用于保持和服务器端的通讯
		//如果服务器有消息推送,则显示在客户端
		go serverHolder(conn)
		//1. 显示登录成功后的菜单. 循环显示
		showMenu()
	} else if loginResMsg.StatusCode == 500 {
		fmt.Println(loginResMsg.Error)
	}
	return
}

