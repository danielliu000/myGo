package model

import (
	"chatRoom/common/message"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/go-redis/redis/v8"
	"strconv"
)
var MyUserDao *UserDao

type UserDao struct {
	rdb *redis.Client
}



func NewUserDao (rdb *redis.Client) (userDao *UserDao) {
	userDao  = &UserDao {
		rdb: rdb,
	}
	return
}
func (userDao *UserDao) GetUserById(id int) (user User, err error){
	//去Redis查询用户
	idStr := strconv.Itoa(id)
	ctx := context.Background()
	val, err := userDao.rdb.HGet(ctx, "users", idStr).Result()
	if err == redis.Nil {
		err = ERROR_USER_NO_EXISTS
	} else if err != nil {
		panic(err)
	}
	//需要将rs反序列化成user对象
	err = json.Unmarshal([]byte(val), &user)
	if err != nil {
		err = errors.New("反序列化失败")
	}
	return
}

func (userDao *UserDao) Register (user *message.User) (err error) {

	_, err = userDao.GetUserById(user.UserId)
	if err == nil {
		err = ERROR_USER_EXISTS
		return
	}

	//有err，则说明查询不到 =》 没有该用户,则该用户可以入库，完成注册
	data, err := json.Marshal(user)
	if err != nil {
		return
	}
	ctx := context.Background()
	val, err := userDao.rdb.HSet(ctx, "users", user.UserId, string(data)).Result()
	if val != 1 || err != nil {
		fmt.Println("注册写入数据错误: ", err)

	} else {
		fmt.Println("注册成功")
	}
	return
}
//完成登录校验
func (userDao *UserDao) Login (userId int, userPwd string) (user User, err error){

	user ,err = userDao.GetUserById(userId)
	if err != nil {
		fmt.Println(err)
		return
	}

	if user.UserPwd != userPwd {
		err = ERROR_USER_PWD
		return
	}
	return

}