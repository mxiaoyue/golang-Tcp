package process

import (
	"demo/goproject/src/go_code/tcp2/client/model"
	"demo/goproject/src/go_code/tcp2/common/message"
	"fmt"
)

// 在用户登入成功后，对CurUser初始化
var CurUser model.CurUser

// 客户端维护的map
var onlineUsers map[int]*message.User = make(map[int]*message.User, 10)

// 在客户端显示当前在线用户
func outputOnlineUser() {

	fmt.Println("当前在线用户列表")
	for id := range onlineUsers {
		fmt.Printf("用户id:%d\n", id)
	}
	fmt.Println()
}

// 更新在线用户
func updateUserStatus(notifyUserStatusMes *message.NotifyUserStatusMes) {
	user, ok := onlineUsers[notifyUserStatusMes.UserId]
	if !ok {
		user = &message.User{
			UserId: notifyUserStatusMes.UserId,
		}
	}
	user.UserStatus = notifyUserStatusMes.Status
	onlineUsers[notifyUserStatusMes.UserId] = user
	outputOnlineUser()
}
