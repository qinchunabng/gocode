package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
	"strings"
)

func main() {
	conn, err := net.Dial("tcp", "127.0.0.1:8888")
	if err != nil {
		fmt.Println("client dial err=", err)
		return
	}
	fmt.Printf("conn=%v\n", conn)
	//功能一：客户端可以发送单行数据，然后就退出
	//os.Stdin代表终端标准输入
	reader := bufio.NewReader(os.Stdin)
	for {
		//从终端读取用户输入并发送给服务器
		line, err := reader.ReadString('\n')
		if err != nil {
			fmt.Println("readString err=", err)
			return
		}
		line = strings.Trim(line, "\r\n")
		if line == "exit" {
			fmt.Println("客户端退出")
			break
		}
		//再将line发给服务器
		_, err = conn.Write([]byte(line))
		if err != nil {
			fmt.Println("conn.Write err=", err)
		}
	}
	//fmt.Printf("客户端发送了%d字节，并退出", n)
}
