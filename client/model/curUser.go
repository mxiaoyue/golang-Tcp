package model

import (
	"demo/goproject/src/go_code/tcp2/common/message"
	"net"
)

// 在客户端，很多地方都会使用到CurUser，所以做成全局
type CurUser struct {
	Conn net.Conn //维护自身的链接
	message.User
}
