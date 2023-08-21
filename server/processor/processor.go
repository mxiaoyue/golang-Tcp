package processor

import (
	"demo/goproject/src/go_code/tcp2/common/message"
	"demo/goproject/src/go_code/tcp2/server/processes"
	"demo/goproject/src/go_code/tcp2/server/utils"
	"fmt"
	"net"
)

/*
总控
用来接收客户端发送来的消息
并进行分发处理
不知道什么原因 该文件放入main目录下会报出undefined: Processor错误
确定定义好了结构体
已经找到原因了 VsCode会报错，cmd下正常运行
目前没找到解决办法
*/
type Processor struct {
	Conn net.Conn
}

// 编写一个SeverProcessMes函数
// 功能：根据客户端发送消息种类不同，决定调用哪个函数来处理
func (p *Processor) serverProcessMes(mes *message.Message) (err error) {

	switch mes.Type {
	case message.LoginMesType:
		//处理登入逻辑
		//创建一个UserProcess实力
		up := &processes.UserProcess{
			Conn: p.Conn,
		}
		err = up.ServerProcessLogin(mes)
	case message.RegisterMesType:
		//处理注册逻辑
		up := &processes.UserProcess{
			Conn: p.Conn,
		}
		err = up.ServerProcessRegister(mes)
	case message.SmsMesType:
		var smsProcess processes.SmsProcess
		smsProcess.SendGroupMes(mes)
	default:
		fmt.Println("消息类型不存在")
	}
	return
}

// 处理和客户端的通讯
func (p *Processor) ProcessFlow() (err error) {
	defer p.Conn.Close()

	//读取客户端发送消息
	for {
		ts := &utils.Transfer{
			Conn: p.Conn,
		}

		mes, err := ts.ReadPkg()
		// if err == io.EOF {
		// 	fmt.Println("客户端退出，服务端也退出")
		// 	return
		// }

		// if err != nil {
		// 	fmt.Println("readPkg() err=", err)
		// 	return
		// }
		if err != nil {
			return err
		}

		err = p.serverProcessMes(&mes)
		if err != nil {
			fmt.Println("serverProcessMes err=", err)
			return err
		}

	}

}
