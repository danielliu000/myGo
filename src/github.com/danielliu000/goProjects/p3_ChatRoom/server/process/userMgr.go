package process

import "fmt"

//UserMgr 在Server只有一个 在很多地方都要用到 因此定义为全局变量
type UserMgr struct {
	onlineUsers map[int]*UserProcess
}
var (
	userMgr *UserMgr
)

//初始化
func init() {
	userMgr = &UserMgr{
		onlineUsers: make(map[int]*UserProcess,1024),
	}
}

func (userMgr *UserMgr) AddOnlineUser(up *UserProcess) {
	userMgr.onlineUsers[up.UserId] = up
}

func (userMgr *UserMgr) DelOnlineUser(userId int) {
	delete(userMgr.onlineUsers, userId)
}

func (userMgr *UserMgr) GetAllOnlineUsers() map[int]*UserProcess {
	return userMgr.onlineUsers
}

func (userMgr *UserMgr) GetOnlineUserById(userId int) (up *UserProcess, err error) {
	up, ok := userMgr.onlineUsers[userId]
	if !ok {//当前用户不在线
		err = fmt.Errorf("用户%d 不存在", userId)
		return
	}
	return
}


