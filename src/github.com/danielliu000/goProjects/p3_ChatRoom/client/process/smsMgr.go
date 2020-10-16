package process

import (
	"chatRoom/common/message"
	"encoding/json"
	"fmt"
)

func OutputGroupMsg(msg *message.Message) {
	//1.反序列化
	var smsMsg message.SmsMsg
	err := json.Unmarshal([]byte(msg.Data), &smsMsg)
	if err != nil {
		fmt.Println("反序列化失败")
		return
	}
	//显示信息
	info := fmt.Sprintf("用户id:\t%d 对大家说:\t%s",
		smsMsg.UserId, smsMsg.Content )
	fmt.Println(info)
	fmt.Println()

}