package process

import (
	"demo/goproject/src/go_code/tcp2/client/utils"
	"demo/goproject/src/go_code/tcp2/common/message"
	"encoding/json"
	"fmt"
	"net"
	"os"
)

// 保持跟服务器端通讯
func ShowMenu() {
	fmt.Printf("============恭喜%d登入成功============\n", CurUser.UserId)
	fmt.Println("1.显示在线用户列表")
	fmt.Println("2.发送消息")
	fmt.Println("3.信息列表")
	fmt.Println("4.退出系统")
	fmt.Println("请选择(1-4)")
	var key int
	var content string
	var smsProcess = &smsProcess{} //空结构体，只是用来调用相关方法
	fmt.Scanf("%d\n", &key)
	switch key {
	case 1:
		outputOnlineUser()
	case 2:
		fmt.Println("你想向大家说点什么")
		fmt.Scanf("%s\n", &content)
		smsProcess.SendGroupMes(content)
	case 3:
		fmt.Println("信息列表")
		for _, v := range infoLog {
			fmt.Print(v)
		}
		fmt.Println()
	case 4:
		fmt.Println("你选择退出系统")
		//Exit(0),程序正常退出。程序会立即终止，任何尚未执行的代码都不会被执行
		os.Exit(0)
	default:
		fmt.Println("输入错误，请重新输入")
	}

}

func serverProcessMes(conn net.Conn) {
	tf := utils.Transfer{
		Conn: conn,
	}
	for {
		fmt.Println("客户端正在等待服务器端消息")
		mes, err := tf.ReadPkg()
		if err != nil {
			fmt.Println("tf.ReadPkg err=", err)
			return
		}
		//读取到则进行下一步逻辑
		switch mes.Type {
		case message.NotifyUserStatusMesType: //有人上线了
			//反序列化返回的消息，并且更新客户端维护的在线用户map
			var notifyUserStatusMes message.NotifyUserStatusMes
			json.Unmarshal([]byte(mes.Data), &notifyUserStatusMes)
			updateUserStatus(&notifyUserStatusMes)
		case message.SmsResMesType: //有人群发消息
			outputGroupMes(&mes)
		default:
			fmt.Println("服务器端返回未知类型")
		}

	}
}
