package process

import (
	"chatRoom/common/message"
	"chatRoom/server/model"
	"chatRoom/server/utils"
	"encoding/json"
	"errors"
	"fmt"
	"net"
)
type UserProcess struct {
	Conn net.Conn
	//添加一个字段，表示Conn是哪个用户的
	UserId int
}
//通知所有在线用户的方法
func(userProcess *UserProcess) NotifyOtherOnlineUser(userId int) {
	//遍历onlineUsers, 然后一个个发送
	for id, up := range userMgr.onlineUsers {
		if id == userId {
			continue
		}
		//开始通知
		up.NotifyOnline(userId)

	}
}
func (userProcess *UserProcess) NotifyOnline(userId int)  {
	//组装NotifyUserStatusType

	notifyUserStatusMsg := message.NotifyUserStatusMsg{
		UserId: userId,
		Status: message.UserOnline,
	}
	//notifyUserStatusMsg序列化
	data, err := json.Marshal(notifyUserStatusMsg)
	if err != nil {
		fmt.Println("notifyUserStatusMsg Marshal error,", err)
		return
	}
	msg := message.Message{
		Type: message.NotifyUserStatusMsgType,
		Data: string(data),
	}
	//对msg序列化
	data, err = json.Marshal(msg)
	if err != nil {
		fmt.Println("msg Marshal error,", err)
		return
	}
	//发送， 创建Transfer实例
	tf := utils.Transfer{
		Conn: userProcess.Conn,
		Buf:  [8096]byte{},
	}
 err =	tf.WritePkg(data)
 if err != nil {
 	fmt.Println("NotifyOnline Error", err)
 	return
 }


}
//编写一个ServerProcessRegister函数 处理注册
func(userProcess *UserProcess) ServerProcessRegister (msg *message.Message) (err error) {
	var registerMsg message.RegisterMsg
	err = json.Unmarshal([]byte(msg.Data), &registerMsg)
	if err != nil {
		err = errors.New("反序列化失败")
	}

	//1. 声明一个RegisterResMsg
	RegisterResMsg := message.RegisterResMsg {
		StatusCode: 0,
		Error:      "",
	}
	//2. 声明一个resMsg
	resMsg := message.Message {  //需要序列化后返回
		Type: message.RegisterResMsgType,
		Data: "",
	}

	//验证登录信息
	//1. 使用 model.MyUserDao 到 Redis中验证
	 err = model.MyUserDao.Register(&registerMsg.User)
	if err != nil {
		if err == model.ERROR_USER_EXISTS {
			RegisterResMsg.StatusCode = 400
			RegisterResMsg.Error = model.ERROR_USER_EXISTS.Error()
		} else {
			RegisterResMsg.StatusCode = 505
			RegisterResMsg.Error = "Unknown Error"
		}
	} else {
		RegisterResMsg.StatusCode = 200
	}
	//3.将loginResMsg 序列化
	data, err := json.Marshal(RegisterResMsg)
	if err != nil {
		err = errors.New("RegisterResMsg 序列化失败")
		return
	}
	//4.将data赋值给resMsg.Data
	resMsg.Data = string(data)
	//5.对resMsg序列化，准备发送
	data, err = json.Marshal(resMsg)
	if err != nil {
		err = errors.New("ResMsg 序列化失败")
		return
	}
	//6.发送data，将发送函数封装到writePkg函数中
	//需要创建utils中的Transfer实例
	tf := &utils.Transfer{
		Conn: userProcess.Conn,
		Buf:  [8096]byte{},
	}
	err = tf.WritePkg(data)
	return
}
//编写一个ServerProcessLogin函数 处理登录
func (userProcess *UserProcess) ServerProcessLogin(msg *message.Message) (err error) {
	//1.从msg中取出msg.Data, 并反序列化成loginMsg
	var loginMsg message.LoginMsg
	err = json.Unmarshal([]byte(msg.Data), &loginMsg)
	if err != nil {
		err = errors.New("反序列化失败")
	}
	//1. 声明一个LoginResMsg
	loginResMsg := message.LoginResMsg {
		StatusCode: 0,
		Error:      "",
	}
	//2. 声明一个resMsg
	resMsg := message.Message {  //需要序列化后返回
		Type: message.LoginResMsgType,
		Data: "",
	}

	//验证登录信息
	//1. 使用 model.MyUserDao 到 Redis中验证
	user, err := model.MyUserDao.Login(loginMsg.UserId, loginMsg.UserPwd)
	if err != nil {
		loginResMsg.StatusCode = 500
	} else {
		loginResMsg.StatusCode = 200
		//将登录的UserId，放入userProcess 以及 loginResMsg中
		//放入userProcess
		//这里 用户登录成功后，将该用户放入 UserMgr 在线用户中
		userProcess.UserId = loginMsg.UserId
		userMgr.AddOnlineUser(userProcess)
		//上线后 通知其他用户状态
		userProcess.NotifyOtherOnlineUser(loginMsg.UserId)

		//放入loginResMsg, 遍历userMgr.onlineUsers
		for id, _ := range userMgr.onlineUsers {
			loginResMsg.UserIds = append(loginResMsg.UserIds, id)
		}


	}
	fmt.Println(user, "登录成功")


	//
	//假设 登录用户名 100 密码123456为合法 则
	//if loginMsg.UserId == 100 && loginMsg.UserPwd == "123456" {
	//	//合法
	//	loginResMsg.StatusCode = 200
	//} else {
	//	//不合法
	//	loginResMsg.StatusCode = 500
	//	loginResMsg.Error = "用户不存在，请注册"
	//}
	//3.将loginResMsg 序列化
	data, err := json.Marshal(loginResMsg)
	if err != nil {
		err = errors.New("loginResMsg 序列化失败")
		return
	}
	//4.将data赋值给resMsg.Data
	resMsg.Data = string(data)
	//5.对resMsg序列化，准备发送
	data, err = json.Marshal(resMsg)
	if err != nil {
		err = errors.New("ResMsg 序列化失败")
		return
	}
	//6.发送data，将发送函数封装到writePkg函数中
	//需要创建utils中的Transfer实例
	tf := &utils.Transfer{
		Conn: userProcess.Conn,
		Buf:  [8096]byte{},
	}
	err = tf.WritePkg(data)
	return
}