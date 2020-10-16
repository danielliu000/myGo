package process

import (
	"chatRoom/client/utils"
	"chatRoom/common/message"
	"encoding/json"
	"fmt"
)

type SmsProcess struct {

}

//发送群聊消息

func (smsProcess *SmsProcess) SendGroupMsg (content string) (err error) {
	var msg message.Message
	msg.Type = message.SmsMsgType

	var smsMsg message.SmsMsg
	smsMsg.Content = content
	smsMsg.UserId = curUser.UserId
	smsMsg.UserStatus = curUser.UserStatus

	data, err := json.Marshal(smsMsg)
	if err != nil {
		fmt.Println("SendGroupMsg 序列化失败，", err)
		return
	}
	msg.Data = string(data)

	data, err = json.Marshal(msg)
	if err != nil {
		fmt.Println("SendGroupMsg 序列化失败，", err)
		return
	}
	tf := utils.Transfer{
		Conn: curUser.Conn,
		Buf:  [8096]byte{},
	}
	err = tf.WritePkg(data)
	if err != nil {
		fmt.Println("SendGroupMsg err=，", err)
		return
	}
	return
}