package process

import (
	"chatRoom/client/utils"
	"chatRoom/common/message"
	"encoding/json"
	"fmt"
	"net"
	"os"
)

//显示登录成功后的界面
func showMenu() {

	for {
		fmt.Println()
		fmt.Println()
		fmt.Println("-------------恭喜xxx登录成功-------------")
		fmt.Println("-------------1. 显示在线用户列表-------------")
		fmt.Println("-------------2. 发送消息-------------")
		fmt.Println("-------------3. 信息列表-------------")
		fmt.Println("-------------4. 退出系统-------------")
		fmt.Println("请选择(1-4): ")

		var key int
		var content string
		smsProcess := &SmsProcess{}
		_, _ = fmt.Scanf("%d\n", &key)
		switch key {
		case 1:
			outputOnlineUser()
		case 2:
			fmt.Println("请输入你想对大家说点什么")
			_, _ = fmt.Scanf("%s\n", &content)
			err := smsProcess.SendGroupMsg(content)

			if err != nil {
				fmt.Println("发送消息失败")
				return
			}
		case 3:
			fmt.Println(3)
		case 4:
			fmt.Println("退出系统")
			os.Exit(0)
		default:
			fmt.Println("你输入的选项不正确")
		}
	}

}

//和服务器保持通讯
func serverHolder(conn net.Conn) {
	//创建Transfer实例, 不停读取服务器消息
	tf := &utils.Transfer{
		Conn: conn,
	}
	for {
		fmt.Println("客户端正在等待服务器发送的消息...")
		msg, err := tf.ReadPkg()
		if err != nil {
			fmt.Println("read error :", err)
			return
		}

		//读到消息
		//fmt.Println(msg)
		switch msg.Type {
		case message.NotifyUserStatusMsgType:
			//有人上线
			//处理
			//1.取出NotifyUserStatusMsg
			notifyUserStatusMsg := message.NotifyUserStatusMsg{}
			err = json.Unmarshal([]byte(msg.Data), &notifyUserStatusMsg)
			if err != nil {
				fmt.Println("反序列化失败，", err)
				return
			}
			//2.将用户加入到客户端map[int]User中
			updateUserStatus(&notifyUserStatusMsg)
		case message.SmsMsgType:
			//有人群发消息
			OutputGroupMsg(&msg)

		default:
			fmt.Println("服务器返回无法处理消息")
		}

	}
}
