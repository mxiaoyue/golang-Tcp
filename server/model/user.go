package model

type User struct {
	//确定字段信息
	//为了序列化和反序列化成功
	//用户信息的json字符串的key 和 结构体字段对应的 tag 名字一致。否则报错
	UserId     int    `json:"userId"`
	UserPwd    string `json:"userPwd"`
	UserName   string `json:"userName"`
	UserStatus int    `json:"userStatus"`
}
