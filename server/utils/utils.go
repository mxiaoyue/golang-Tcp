package utils

import (
	"demo/goproject/src/go_code/tcp2/common/message"
	"encoding/binary"
	"encoding/json"
	"fmt"
	"net"
)

// 这里将这些方法关联到结构体中
type Transfer struct {
	//分析它应该有哪些字段
	Conn net.Conn
	Buf  [8096]byte //这是传输时使用的缓冲
}

// 将message序列化好后的[]byte发送给客户端
func (tf *Transfer) WritePkg(data []byte) (err error) {

	var pkgLen = uint32(len(data))
	//一个字节有8位，所以4个字符就是32位字节，用来保存unit32的数据
	// var buf = make([]byte, 4)
	//binary.BigEndian.PutUint32([]byte,unit32)	这个方法将一个unit32的数放入一个切片内
	binary.BigEndian.PutUint32(tf.Buf[0:4], pkgLen)
	//net包内的Write只能发送byte切片。需要事先发送长度切片进行比较确认数据有无丢失
	_, err = tf.Conn.Write(tf.Buf[:4])
	//n!=4是因为Write返回值是发送了多少字符
	if err != nil {
		fmt.Println("conn.Write err=", err)
		return
	}

	n, err := tf.Conn.Write(data)
	if n != int(pkgLen) || err != nil {
		fmt.Println("conn.Write err=", err)
		return
	}

	return
}

// 读取数据包封装成一个函数readPkg(),返回Message,err
func (tf *Transfer) ReadPkg() (mes message.Message, err error) {
	// buf := make([]byte, 8096)
	fmt.Println("等待读取客户端发送的数据")
	//conn.Read在conn没有关闭的情况下，才会阻塞
	//如果客户端关闭了conn则，不会阻塞
	_, err = tf.Conn.Read(tf.Buf[:4])
	if err != nil {
		// err = errors.New("read pkg heade error")
		return
	}
	// fmt.Println("读到的长度为=", buf[0:4])
	//根据buf[:4]转换成uint32类型
	//该方法将一个[]byte转换为uint32
	var pkgLen = binary.BigEndian.Uint32(tf.Buf[:4])
	n, err := tf.Conn.Read(tf.Buf[:pkgLen])
	if uint32(n) != pkgLen || err != nil {
		// fmt.Println("conn.Read err=", err)
		// err = errors.New("read pkg body error")
		return
	}
	//把pkg反序列化成->message.Message
	//这里需要加上&符，不然反序列出来是空的
	err = json.Unmarshal(tf.Buf[:pkgLen], &mes)
	if err != nil {
		fmt.Println("json.Unmarshal err=", err)
		return
	}
	return

}
