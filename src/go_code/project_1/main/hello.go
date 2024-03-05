package main

import (
	"fmt"
	"net"
)

func process(conn net.Conn) {
	defer conn.Close()
	for {
		buf := make([]byte, 1024)
		n, err := conn.Read(buf)
		if err != nil {
			return
		}
		fmt.Println(string(buf[0:n]))
	}
}

func main() {
	//fmt.Println("hello word")
	fmt.Println("服务端启动了....")
	listen, err := net.Listen("tcp", "127.0.0.1:8080")
	if err != nil {
		fmt.Println("监听失败,err:", err)
		return
	}

	for {
		conn, err2 := listen.Accept()
		if err2 != nil {
			fmt.Println("客户端的等待失败，err2:", err2)
		} else {
			fmt.Println("等待连接成功，con=%v,接收客户端消息：%v", conn, conn.RemoteAddr().String())
		}
		go process(conn)
	}
}
