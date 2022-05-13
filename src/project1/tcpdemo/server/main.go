package main

import (
	"fmt"
	"net"
)

func main() {
	fmt.Println("服务器开始监听...")
	//net.Listen("tcp", "0.0.0.0:8888")
	//1. tcp表示网络协议是tcp
	//2. 0.0.0.0:8888表示本地监听8888端口
	listen, err := net.Listen("tcp", "0.0.0.0:8888")
	if err != nil {
		fmt.Println("listen err=", err)
		return
	}
	defer listen.Close()

	//等待客户端连接
	for {
		fmt.Println("等待客户端连接...")
		conn, err := listen.Accept()
		if err != nil {
			fmt.Println("Accept() err=", err)
		} else {
			fmt.Printf("Accept() suc con=%v, client ip=%v\n", conn, conn.RemoteAddr().String())
		}
		go process(conn)
	}
}

func process(conn net.Conn) {
	defer conn.Close()

	for {
		buf := make([]byte, 1024)
		//等待客户端发送，如果客户端没有Write
		//就会超时阻塞
		n, err := conn.Read(buf)
		if err != nil {
			fmt.Println("客户端退出")
			//fmt.Println("服务端的read err=", err)
			return
		}
		//显示客户端发送的内容到服务器的终端
		fmt.Println(string(buf[:n]))
	}
}
