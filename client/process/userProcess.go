package process

import (
	"demo/goproject/src/go_code/tcp2/client/utils"
	"demo/goproject/src/go_code/tcp2/common/message"
	"encoding/json"
	"fmt"
	"net"
	"os"
)

type UserProcess struct {
	Mes message.Message
}

func (u *UserProcess) Register(userId int, userPwd, userName string) {
	//建立连接
	conn, err := net.Dial("tcp", "localhost:8889")
	if err != nil {
		fmt.Println("net.Dial err=", err)
		return
	}
	defer conn.Close()
	//创建mes协议，协议头为注册账户
	var mes message.Message
	mes.Type = message.RegisterMesType
	//接受实参创建实例并序列化，并且赋值给mes的data字段
	var registerMes message.RegisterMes
	registerMes.User.UserId = userId
	registerMes.User.UserPwd = userPwd
	registerMes.User.UserName = userName
	data, err := json.Marshal(registerMes)
	if err != nil {
		fmt.Println("json.Marshal err=", err)
		return
	}
	mes.Data = string(data)
	//将协议序列化发送到服务器
	data, err = json.Marshal(mes)
	if err != nil {
		fmt.Println("json.Marshal err=", err)
		return
	}
	tf := &utils.Transfer{
		Conn: conn,
	}
	tf.WritePkg(data)
	//接受服务器返回的协议
	mes, err = tf.ReadPkg()
	if err != nil {
		fmt.Println("readPkg err=", err)
	}
	// 将mes的data部分反序列化成RegisterResMes
	var registerResMes message.RegisterResMes
	err = json.Unmarshal([]byte(mes.Data), &registerResMes)
	if err != nil {
		fmt.Println("json.unmarshal err=", err)
	}
	//识别状态码
	if registerResMes.Code == 200 {
		fmt.Println("注册成功，请重新登入")
	} else {
		fmt.Println(registerResMes.Error)
	}

	//结束
	os.Exit(0)
}

// 登入函数
func (u *UserProcess) Login(userId int, userPassword string) (err error) {
	//连接到服务器端,这边之后要改成配置文件
	conn, err := net.Dial("tcp", "localhost:8889")
	if err != nil {
		fmt.Println("net.Dial err=", err)
		return
	}
	defer conn.Close()
	//通过conn发送消息给服务器
	var mes message.Message
	//设置发送类型
	mes.Type = message.LoginMesType
	//创建一个loginMessage 结构体
	var loginMes message.LoginMes
	loginMes.UserId = userId
	loginMes.UserPwd = userPassword
	//将loginMessage序列化,放入mes内。直接放入再序列化的话 会将指针序列化，而不是想要的数据序列化
	data, err := json.Marshal(loginMes)
	if err != nil {
		fmt.Println("json.Marshal err=", err)
		return
	}
	//赋值
	mes.Data = string(data)
	//将mes再次序列化。进行发送
	data, err = json.Marshal(mes)
	if err != nil {
		fmt.Println("json.Marshal err=", err)
		return
	}
	//调用发送方法
	tf := &utils.Transfer{
		Conn: conn,
	}
	tf.WritePkg(data)
	//处理服务器端返回的消息
	mes, err = tf.ReadPkg()
	if err != nil {
		fmt.Println("readPkg err=", err)
	}
	//将mes的data部分反序列化成loginResMes
	var loginResMes message.LoginResMes
	err = json.Unmarshal([]byte(mes.Data), &loginResMes)
	if err != nil {
		fmt.Println("json.unmarshal err=", err)
	}
	//登入成功时
	if loginResMes.Code == 200 {
		//初始化curUser
		CurUser.Conn = conn
		CurUser.UserId = loginMes.UserId
		CurUser.UserStatus = message.UserOnline
		//将所有在线用户维护到onlineUsers map中
		for _, v := range loginResMes.UsersId {
			if v == userId {
				continue
			}
			user := &message.User{
				UserId:     v,
				UserStatus: message.UserOnline,
			}
			onlineUsers[v] = user
		}

		//需要启动一个协程，该协程保持与服务器端的通讯。如果服务器有数据推送给客户端，则接收并显示
		go serverProcessMes(conn)
		//显示登入成功过后的菜单
		for {
			ShowMenu()
		}
	} else {
		fmt.Println(loginResMes.Error)
	}
	return nil
}
