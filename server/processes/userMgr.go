package processes

import "fmt"

// 因为UserMgr实例在服务器有且只有一个
// 因为很多地方都会使用到，因此作为一个全局变量
var (
	userMgr *UserMgr
)

type UserMgr struct {
	//int为用户的ID，UserProcess为用户的conn链接
	onlineUsers map[int]*UserProcess
}

func init() {
	userMgr = &UserMgr{
		onlineUsers: make(map[int]*UserProcess, 1024),
	}
}

// 完成对onlineUsers的添加/修改
func (um *UserMgr) AddOnlineUser(up *UserProcess) {
	um.onlineUsers[up.UserId] = up
}

// 删除
func (um *UserMgr) DelOnlineUser(userId int) {
	delete(um.onlineUsers, userId)
}

// 返回所有在线用户
func (um *UserMgr) GetAllOnlineUser() (onlineUsers map[int]*UserProcess) {
	return um.onlineUsers
}

// 根据ID返回用户
func (um *UserMgr) GetOnlineUserById(userId int) (up *UserProcess, err error) {
	up, ok := um.onlineUsers[userId]
	if !ok { //说明当前用户不在线
		err = fmt.Errorf("用户%d不在线", userId)
	}
	return
}
