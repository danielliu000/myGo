package processor

import (
	"chatRoom/common/message"
	"chatRoom/server/process"
	"chatRoom/server/utils"
	"fmt"
	"io"
	"net"
)
type Processor struct {
	Conn net.Conn
}
//编写一个ServerProcessMsg函数
//功能：根据客户端发送消息种类不同，决定调用哪个函数来处理
func (processor *Processor) serverProcessMsg(msg *message.Message) (err error) {

	//看看是否能接收到客户端发送的群发消息
	fmt.Println("msg = ", msg)
	switch msg.Type {
	case message.LoginMsgType:
		up := &process.UserProcess{Conn:processor.Conn}
		err = up.ServerProcessLogin(msg)
	case message.RegisterMsgType:
	//处理注册信息
		up := &process.UserProcess{Conn:processor.Conn}
		err = up.ServerProcessRegister(msg)
	case message.SmsMsgType:
		sp := &process.SmsProcess{}
		err = sp.SendGroupMsg(msg)
	default:
		fmt.Println("消息类型不存在，无法处理")
	}
	return
}
func (processor *Processor) ProcessHandler() (err error) {
	for {
		//读取数据包
		tf := &utils.Transfer{
			Conn: processor.Conn,
			Buf:  [8096]byte{},
		}
		msg, err := tf.ReadPkg()
		if err != nil {
			if err == io.EOF {
				break
			}
			fmt.Println(err)
			return err
		}
		err = processor.serverProcessMsg(&msg)
		if err != nil {
			fmt.Println("服务器处理失败:", err)
			return err
		}
	}
	return err
}
