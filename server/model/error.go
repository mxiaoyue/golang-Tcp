package model

import "errors"

var (
	ErrUserNotExists = errors.New("用户不存在")
	ErrUserEexists   = errors.New("用户已存在")
	ErrUserPwd       = errors.New("密码不正确")
)
