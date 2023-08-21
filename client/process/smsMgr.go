package process

import (
	"demo/goproject/src/go_code/tcp2/common/message"
	"encoding/json"
	"fmt"
	"time"
)

var infoLog []string

func outputGroupMes(mes *message.Message) {
	var smsResMes message.SmsResMes
	err := json.Unmarshal([]byte(mes.Data), &smsResMes)
	if err != nil {
		fmt.Println("json.unmarshal err=", err.Error())
		return
	}

	info := fmt.Sprintf("用户id:\t%d 对大家说:\t%s\n", smsResMes.UserId, smsResMes.Content)
	fmt.Print(info)
	fmt.Println()
	addInfoLog(info)
}

func addInfoLog(str string) {
	currentTime := time.Now()
	currentTimeString := currentTime.Format("2006-01-02 15:04:05")
	str = currentTimeString + str
	infoLog = append(infoLog, str)
}
