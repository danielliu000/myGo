package model

import "errors"

var (
	ERROR_USER_NO_EXISTS = errors.New("用户不存在")
	ERROR_USER_EXISTS = errors.New("用户已存在")
	ERROR_USER_PWD = errors.New("用户密码错误")

		)