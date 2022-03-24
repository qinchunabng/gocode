package main

import (
	"flag"
	"fmt"
)

func main() {
	var user string
	var pwd string
	var host string
	var port int
	//获取命令行调用参数
	//参数-u,默认为空，赋值给user
	flag.StringVar(&user, "u", "", "用户名，默认为空")
	flag.StringVar(&pwd, "p", "", "密码，默认为空")
	flag.StringVar(&host, "h", "localhost", "主机名，默认为localhost")
	flag.IntVar(&port, "P", 3306, "端口号，默认3306")

	flag.Parse()

	fmt.Printf("user=%v,pwd=%v,host=%v,port=%v\n", user, pwd, host, port)
}
