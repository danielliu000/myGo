package utils

import (
	"chatRoom/common/message"
	"encoding/binary"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net"
)
//这里将方法关联到结构体中
type Transfer struct {
	Conn net.Conn
	Buf [8096]byte  //传输时使用的缓冲
}

func (transfer *Transfer) ReadPkg() (msg message.Message, err error) {
	//buf := make([]byte, 8096)
	//这里读取数据包，直接封装成一个函数readPkg(conn),返回 message, error
	_, err = transfer.Conn.Read(transfer.Buf[:4])

	if err != nil {
		if err == io.EOF {
			fmt.Println("客户端退出连接")
			return
		}
		fmt.Println("conn.Read error=", err)
		return
	}
	//根据buf[:4] 转换成一个uint32类型
	pkgLen := binary.BigEndian.Uint32(transfer.Buf[:4])

	//根据pkgLen读取消息内容
	n, err := transfer.Conn.Read(transfer.Buf[:pkgLen]) //从conn里读取pkgLen个字节，扔到buf中
	if n != int(pkgLen) || err != nil {
		err = errors.New("read pkg head error")
		return
	}

	//将pkgLen反序列化成 message.Message
	err = json.Unmarshal(transfer.Buf[:pkgLen], &msg)
	if err != nil {
		fmt.Println("Json.Unmarshal failed: ", err)
		return
	}

	return
}
func (transfer *Transfer) WritePkg(data []byte) (err error) {
	//先发送长度
	pkgLen := uint32(len(data))
	binary.BigEndian.PutUint32(transfer.Buf[0:4], pkgLen)

	n, err := transfer.Conn.Write(transfer.Buf[:4])
	if n != 4 || err != nil {
		fmt.Println("conn.Write(bytes) fail, ", err)
		return
	}

	//发送数据本身
	n, err = transfer.Conn.Write(data)
	if n != int(pkgLen) || err != nil {
		fmt.Println("conn.Write(bytes) fail, ", err)
		return
	}
	return

}
