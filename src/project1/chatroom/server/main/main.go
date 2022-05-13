package main

import (
	"fmt"
	"net"
	"project1/chatroom/server/model"
	"time"
)

func main() {
	//初始化redis的连接池
	initPool("localhost:6379", 16, 0, 300*time.Second)
	initUserDao()
	fmt.Println("服务器在8889端口监听...")
	listen, err := net.Listen("tcp", "0.0.0.0:8889")
	if err != nil {
		fmt.Println("net.Listen err=", err)
		return
	}
	defer listen.Close()

	//等待客户端连接
	for {
		fmt.Println("等待客户端连接服务器...")
		conn, err := listen.Accept()
		if err != nil {
			fmt.Println("listen.Accept err=", err)
		}
		//一旦连接成功，则启动一个协程和客户端保持通讯
		go process(conn)
	}
}

//完成对UserDao初始化
func initUserDao() {
	model.MyUserDao = model.NewUserDao(pool)
}

func process(conn net.Conn) {
	defer conn.Close()
	pr := &Processor{
		Conn: conn,
	}
	err := pr.Process()
	if err != nil {
		fmt.Println("客户端和服务器通讯协程出问题了，err=", err)
	}
}
