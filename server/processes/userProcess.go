package processes

import (
	"demo/goproject/src/go_code/tcp2/common/message"
	"demo/goproject/src/go_code/tcp2/server/model"
	"demo/goproject/src/go_code/tcp2/server/utils"
	"encoding/json"
	"fmt"
	"net"
)

type UserProcess struct {
	Conn   net.Conn
	UserId int
}

// 通知所有在线用户的方法
func (u *UserProcess) NotifyOthersOnlineUser(userId int) {
	for id, up := range userMgr.onlineUsers {
		if id == userId {
			continue
		}
		u.NotifyMeOnline(u.UserId, up)
	}
}
func (u *UserProcess) NotifyMeOnline(id int, up *UserProcess) {
	var resMes message.Message
	resMes.Type = message.NotifyUserStatusMesType
	var nus message.NotifyUserStatusMes
	nus.UserId = id
	nus.Status = message.UserOnline
	data, err := json.Marshal(nus)
	if err != nil {
		fmt.Println("NotifyMeOnline() json.Marshal err=", err)
	}
	resMes.Data = string(data)
	data, err = json.Marshal(resMes)
	if err != nil {
		fmt.Println("NotifyMeOnline() json.Marshal err=", err)
	}
	tf := &utils.Transfer{
		Conn: up.Conn,
	}
	tf.WritePkg(data)
}

// 处理用户相关的文件
// 用户注册函数
func (u *UserProcess) ServerProcessRegister(mes *message.Message) (err error) {
	//取出data并反序列化为RegisterMes
	var registerMes message.RegisterMes
	err = json.Unmarshal([]byte(mes.Data), &registerMes)
	if err != nil {
		fmt.Println("json.unmarshal err=", err)
	}
	//注册用户返回协议
	var resMes message.Message
	resMes.Type = message.RegisterResMesType
	//返回的结构体
	var registerResMes message.RegisterResMes
	//调用注册函数，并设置状态码

	err = model.MyUserDao.Register(&registerMes.User)
	if err != nil {
		if err == model.ErrUserEexists {
			registerResMes.Code = 400
			registerResMes.Error = "用户已存在"
		} else {
			registerResMes.Code = 506
			registerResMes.Error = "未知错误"
		}
	} else {
		registerResMes.Code = 200
	}
	//序列化返回协议的参数，并赋值
	data, err := json.Marshal(registerResMes)
	if err != nil {
		fmt.Println("json.Marshal err=", err)
		return
	}
	resMes.Data = string(data)
	data, err = json.Marshal(resMes)
	if err != nil {
		fmt.Println("json.Marshal err", err)
		return
	}
	//调用发送相关函数，将序列化好的内容发送给客户端
	tf := &utils.Transfer{
		Conn: u.Conn,
	}
	tf.WritePkg(data)

	return
}

// 用户登入函数
func (u *UserProcess) ServerProcessLogin(mes *message.Message) (err error) {
	//从mes中取出data，并反序列化成LoginMes
	var loginMes message.LoginMes
	err = json.Unmarshal([]byte(mes.Data), &loginMes)
	if err != nil {
		fmt.Println("json.unmarshal err=", err)
	}
	var resMes message.Message
	resMes.Type = message.LoginResMesType
	//再声明一个LoginResMes
	var loginResMes message.LoginResMes

	//这里暂时先用固定值判断，之后添加数据库验证逻辑
	// if loginMes.UserId == 100 && loginMes.UserPwd == "123456" {
	// //合法
	// loginResMes.Code = 200
	// } else {
	// //不合法
	// loginResMes.Code = 500 //500表示，该用户不存在
	// loginResMes.Error = "该用户不存在，请注册后再使用"
	// }

	//进行redis数据验证，这里的user存储了该用户的所有信息
	user, err := model.MyUserDao.Login(loginMes.UserId, loginMes.UserPwd)
	if err != nil {
		//先用固定的错误测试，可以用switch扩展
		if err == model.ErrUserNotExists {
			loginResMes.Code = 500
			loginResMes.Error = "该用户不存在，请注册后再使用"
		} else if err == model.ErrUserPwd {
			loginResMes.Code = 403
			loginResMes.Error = err.Error()
		} else {
			loginResMes.Code = 505
			loginResMes.Error = "服务器内部错误"
		}

	} else {
		loginResMes.Code = 200
		//登入成功后，添加到在线用户列表
		u.UserId = loginMes.UserId
		userMgr.AddOnlineUser(u)
		//将当前在线用户的id放入到loginResMes.UsersId
		for id := range userMgr.onlineUsers {
			if id == u.UserId {
				continue
			}
			loginResMes.UsersId = append(loginResMes.UsersId, id)
		}
		//通知其他用户我上线了
		u.NotifyOthersOnlineUser(u.UserId)
		fmt.Printf("%v登入成功\n", user)
	}

	//序列化loginResMes并放入resMes内。再次序列化resMes
	data, err := json.Marshal(loginResMes)
	if err != nil {
		fmt.Println("json.Marshal err=", err)
		return
	}

	resMes.Data = string(data)
	data, err = json.Marshal(resMes)
	if err != nil {
		fmt.Println("json.Marshal err", err)
		return
	}
	//调用发送相关函数，将序列化好的内容发送给客户端
	tf := &utils.Transfer{
		Conn: u.Conn,
	}
	tf.WritePkg(data)
	return
}
