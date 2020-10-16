package message

const (
	LoginMsgType            = "LoginMsg"
	LoginResMsgType         = "LoginResMsg"
	RegisterMsgType         = "RegisterMsg"
	RegisterResMsgType      = "RegisterResMsg"
	NotifyUserStatusMsgType = "NotifyUserStatusMsg"
	SmsMsgType              = "SmsMsg"
)

type Message struct {
	Type string `json:"type"` //消息类型
	Data string `json:"data"` //消息内容
}

type LoginMsg struct {
	UserId   int    `json:"user_id"`   //用户ID
	UserPwd  string `json:"user_pwd"`  //用户密码
	UserName string `json:"user_name"` //用户名
}

type LoginResMsg struct {
	StatusCode int    `json:"status_code"` //返回状态码 500---未注册用户 200 --- 登录成功
	Error      string `json:"error"`       // 返回错误信息
	UserIds    []int  //保存用户ID切片
}

type RegisterMsg struct {
	User User `json:"user"`
}
type RegisterResMsg struct {
	StatusCode int    `json:"status_code"`
	Error      string `json:"error"`
}

type NotifyUserStatusMsg struct {
	UserId int `json:"user_id"`
	Status int `json:"status"`
}

//用户状态常量

const (
	UserOnline = iota
	UserOffline
	UserBusyStatus
)

//增加一个smsMsg //发送的消息

type SmsMsg struct {
	Content string `json:"content"`
	User
}
