package process

import (
	"chatRoom/common/message"
	"chatRoom/server/utils"
	"encoding/json"
	"fmt"
	"net"
)

type SmsProcess struct {

}

//写方法转发消息

func(smsProcess *SmsProcess) SendGroupMsg(msg *message.Message) (err error) {
	//遍历服务器端的map
	//将消息转发取出
	//取出msg内容SmsMsg
	var smsMsg message.SmsMsg
	err = json.Unmarshal([]byte(msg.Data), &smsMsg)
	if err != nil {
		fmt.Println("反序列化失败,", err)
		return
	}
	data, err := json.Marshal(msg)
	if err != nil {
		fmt.Println("序列化失败, err=", err)
		return
	}
	for id, up := range userMgr.onlineUsers {
		if id == smsMsg.UserId {
			continue
		}
		smsProcess.SendMsgToEachOnlineUser(data, up.Conn)
	}
	return
}

func (smsProcess *SmsProcess) SendMsgToEachOnlineUser(data []byte, conn net.Conn) {
	//创建Transfer
	tf := &utils.Transfer{
		Conn: conn,
		Buf:  [8096]byte{},
	}
	 err := tf.WritePkg(data)
	 if err != nil {
	 	fmt.Println("转发消息失败，err=", err)
	 }
}