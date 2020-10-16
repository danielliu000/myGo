package process

import (
	"chatRoom/client/model"
	"chatRoom/common/message"
	"fmt"
)

var onlineUsers map[int]*message.User = make(map[int]*message.User, 10)
var curUser model.CurUser //在用户登录成功后，完成对CurUser初始化

//编写方法 处理返回的NotifyUserStatusMsg

func updateUserStatus(notifyUserStatusMsg *message.NotifyUserStatusMsg) {

	user, ok := onlineUsers[notifyUserStatusMsg.UserId]
	if !ok {
		user = &message.User{
			UserId: notifyUserStatusMsg.UserId,
		}
	}
	user.UserStatus = notifyUserStatusMsg.Status
	onlineUsers[notifyUserStatusMsg.UserId] = user
	outputOnlineUser()
}

//显示当前在线用户
func outputOnlineUser () {
	//遍历onlineUser
	fmt.Println("当前在线用户列表：")
	for id, _ := range onlineUsers {

		fmt.Println("用户id:\t", id)
	}
}