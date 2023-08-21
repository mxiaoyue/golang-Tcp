package main

import (
	"demo/goproject/src/go_code/tcp2/server/processor"
	"fmt"
	"io"
	"net"
	"time"
)

// 处理和客户端的通讯
func abcc(conn net.Conn) {
	defer conn.Close()
	pr := &processor.Processor{
		Conn: conn,
	}

	err := pr.ProcessFlow()
	if err == io.EOF {
		fmt.Println("客户端退出，服务端也退出")
		return
	}
	if err != nil {
		fmt.Println("readPkg() err=", err)
		return
	}

}

// 主程序
func main() {
	//服务器启动时，初始化redis连接池
	processor.InitPool("localhost:6379", 16, 0, 300*time.Second)
	processor.InitUserDao()
	listen, err := net.Listen("tcp", "127.0.0.1:8889")
	if err != nil {
		fmt.Println("net.Listen err=", err)
		return
	}
	defer listen.Close()
	//监听成功等待客户端连接
	for {
		fmt.Println("服务器在8889端口监听")
		conn, err := listen.Accept() //这里进行阻塞
		if err != nil {
			fmt.Println("listen.Accept err=", err)
		}
		//一旦连接成功，则启动一个协程与客户端保持通讯

		go abcc(conn)
	}

}
