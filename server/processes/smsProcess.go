package processes

import (
	"demo/goproject/src/go_code/tcp2/common/message"
	"demo/goproject/src/go_code/tcp2/server/utils"
	"encoding/json"
	"fmt"
	"net"
)

type SmsProcess struct {
}

func (sp *SmsProcess) SendGroupMes(mes *message.Message) {
	var smsMes message.SmsMes
	json.Unmarshal([]byte(mes.Data), &smsMes)

	mes.Type = message.SmsResMesType
	data, err := json.Marshal(mes)
	if err != nil {
		fmt.Println("json.Marshal err=", err)
	}

	for id, up := range userMgr.onlineUsers {
		if id == smsMes.UserId {
			continue
		}
		sp.SendMesToEachOnlineUser(data, up.Conn)
	}
}

func (sp *SmsProcess) SendMesToEachOnlineUser(data []byte, conn net.Conn) {

	tf := &utils.Transfer{
		Conn: conn,
	}
	err := tf.WritePkg(data)
	if err != nil {
		fmt.Println("转发消息失败")
	}
}
