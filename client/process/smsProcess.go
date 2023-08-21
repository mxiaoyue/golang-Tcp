package process

import (
	"demo/goproject/src/go_code/tcp2/client/utils"
	"demo/goproject/src/go_code/tcp2/common/message"
	"encoding/json"
	"fmt"
)

type smsProcess struct {
}

// 将消息发送到群聊
func (sp *smsProcess) SendGroupMes(content string) (err error) {
	var mes message.Message
	mes.Type = message.SmsMesType

	var smsMes message.SmsMes
	smsMes.Content = content               //内容
	smsMes.UserId = CurUser.UserId         //发送方Id
	smsMes.UserStatus = CurUser.UserStatus //状态

	data, err := json.Marshal(smsMes)
	if err != nil {
		fmt.Println("json.Marshal err=", err)
		return
	}
	mes.Data = string(data)
	data, err = json.Marshal(mes)
	if err != nil {
		fmt.Println("json.Marshal err=", err)
		return
	}

	tf := &utils.Transfer{
		Conn: CurUser.Conn,
	}
	err = tf.WritePkg(data)
	if err != nil {
		fmt.Println("SendGroupMes err=", err.Error())
		return
	}
	return
}
