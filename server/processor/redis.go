package processor

import (
	"demo/goproject/src/go_code/tcp2/server/model"
	"time"

	"github.com/gomodule/redigo/redis"
)

var Pool *redis.Pool

func InitPool(address string, maxIdle, maxActive int, idleTimeout time.Duration) {
	Pool = &redis.Pool{
		MaxIdle:     maxIdle,     //最大空闲数(连接数)
		MaxActive:   maxActive,   //表示和数据库的最大连接数，0表示不限制
		IdleTimeout: idleTimeout, //最大空闲时间
		Dial: func() (redis.Conn, error) { //初始化连接代码，连接哪个ip的redis
			return redis.Dial("tcp", address)
		},
	}
}

func InitUserDao() {
	//这里的pool是main包下redis.go的全局变量
	//因为需要用到pool 所以等pool初始化后再调用该函数
	model.NewUserDao(Pool)
	// model.MyUserDao = model.NewUserDao(pool)
}
