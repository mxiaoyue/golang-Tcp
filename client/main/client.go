package main

import (
	"demo/goproject/src/go_code/tcp2/client/process"
	"fmt"
)

// 一个表示用户ID 一个表示用户密码
var userId int
var userPassword string
var userName string

func main() {
	//接收用户选择
	var key int
	//判断是否显示菜单
	// var loop = true

	for {
		fmt.Println("============欢迎登入多人聊天系统============")
		fmt.Println("\t\t\t1 登入聊天室")
		fmt.Println("\t\t\t2 注册用户")
		fmt.Println("\t\t\t3 退出系统")
		fmt.Println("请选择(1-3):")

		fmt.Scanf("%d\n", &key)
		switch key {
		case 1:
			fmt.Println("登入聊天室")
			fmt.Println("请输入用户ID:")
			fmt.Scanf("%d\n", &userId)
			fmt.Println("请输入用户密码:")
			fmt.Scanf("%s\n", &userPassword)
			up := &process.UserProcess{}
			up.Login(userId, userPassword)

			// loop = false
		case 2:
			fmt.Println("注册用户")
			fmt.Println("请输入用户ID:")
			fmt.Scanf("%d\n", &userId)
			fmt.Println("请输入用户密码:")
			fmt.Scanf("%s\n", &userPassword)
			fmt.Println("请输入用户昵称:")
			fmt.Scanf("%s\n", &userName)
			//调用UserProcess完成注册
			up := &process.UserProcess{}
			up.Register(userId, userPassword, userName)
			// loop = false
		case 3:
			fmt.Println("退出系统")
			// loop = false
		default:
			fmt.Println("输入有误，请重新输入")
		}
	}

	//判断用户输入，显示新的提示信息
	// if key == 1 {
	// fmt.Println("请输入用户ID:")
	// fmt.Scanf("%d\n", &userId)
	// fmt.Println("请输入用户密码:")
	// fmt.Scanf("%s\n", &userPassword)
	//
	// login(userId, userPassword)
	//
	// } else if key == 2 {
	// fmt.Println("注册逻辑")
	// }
}
