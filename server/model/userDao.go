package model

import (
	"demo/goproject/src/go_code/tcp2/common/message"
	"encoding/json"
	"fmt"

	"github.com/gomodule/redigo/redis"
)

var (
	MyUserDao *UserDao
)

type UserDao struct {
	pool *redis.Pool
}

// 使用工厂模式，创建一个UserDao的实例
func NewUserDao(poola *redis.Pool) (userDao *UserDao) {
	userDao = &UserDao{
		pool: poola,
	}
	MyUserDao = userDao
	return
}

// 根据用户ID返回一个User实例
func (u *UserDao) getUserById(conn redis.Conn, userId int) (user *User, err error) {
	//通过给定的ID去Redis查询用户
	res, err := redis.String(conn.Do("HGet", "users", userId))
	if err != nil {
		if err == redis.ErrNil { //在users这个hash表中没有找到对应id
			err = ErrUserNotExists
		}
		return
	}
	//把res反序列化成实例对象
	err = json.Unmarshal([]byte(res), &user)
	if err != nil {
		fmt.Println("json.Unmarshal err=", err)
		return
	}
	return
}

// 完成登入校验
// Login 完成对用户的验证。
// 如果用户的id和pwd都正确，则返回一个user实例
// 如果用户的id或pwd有错误，则返回对应的错误信息
func (u *UserDao) Login(userId int, userPwd string) (uesr *User, err error) {
	conn := u.pool.Get()
	defer conn.Close()
	uesr, err = u.getUserById(conn, userId)
	if err != nil {
		return
	}
	//这里判断用户密码是否正确
	if uesr.UserPwd != userPwd {
		err = ErrUserPwd
		return
	}
	return
}

// 注册校验
func (u *UserDao) Register(user *message.User) (err error) {
	conn := u.pool.Get()
	defer conn.Close()
	_, err = u.getUserById(conn, user.UserId)
	//nil表示取出数据了，用户已存在
	if err == nil {
		err = ErrUserEexists
		return
	}
	//用户不存在的逻辑
	data, err := json.Marshal(user)
	if err != nil {
		fmt.Println("json.Marshal err=", err)
		return
	}
	//入库
	_, err = conn.Do("HSet", "users", user.UserId, string(data))
	if err != nil {
		fmt.Println("保存注册用户错误 err=", err)
		return
	}
	return

}
