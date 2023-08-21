package message

// 消息类型常量
const (
	LoginMesType            = "LoginMes"            //登入
	LoginResMesType         = "LoginResMes"         //登入返回
	RegisterMesType         = "RegisterMes"         //消息
	RegisterResMesType      = "RegisterResMes"      //消息返回
	NotifyUserStatusMesType = "NotifyUserStatusMes" //通知用户状态
	SmsMesType              = "SmsMes"              //消息
	SmsResMesType           = "SmsResMes"           //消息
)

// 用户在线常量
const (
	UserOnline = iota
	UserOffline
	UserBusyStatus
)

type Message struct {
	Type string `json:"type"` //消息类型
	Data string `json:"data"` //消息内容
}

// 登入相关
type LoginMes struct {
	UserId   int    `json:"userId"`   //用户ID
	UserPwd  string `json:"userPwd"`  //用户密码
	UserName string `json:"userName"` //用户名
}
type LoginResMes struct {
	Code    int    `json:"code"`    //返回状态码 500表示未注册 200表示登入成功
	UsersId []int  `json:"usersId"` //返回所有在线用户的id切片
	Error   string `json:"error"`   //返回错误信息
}

// 注册相关
type RegisterMes struct {
	User User `json:"user"` //User结构体
}
type RegisterResMes struct {
	Code  int    `json:"code"` //返回状态码 400表示已被注册 200表示注册成功
	Error string `json:"error"`
}

// 服务器主动推送用户状态变化的消息
type NotifyUserStatusMes struct {
	UserId int `json:"userId"` //用户ID
	Status int `json:"status"` //状态
}

type SmsMes struct {
	Content string `json:"content"` //内容
	User           //匿名结构体，继承
}
type SmsResMes struct {
	Content string `json:"content"` //内容
	User           //匿名结构体，继承
}
