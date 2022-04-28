package main

import (
	"fmt"
	"time"
)

var userId int
var userPwd string

func main() {
	//接收用户的选择
	var key int
	//是否继续显示菜单
	var loop = true

	for {
		fmt.Println("--------------欢迎登陆多人聊天系统--------------")
		fmt.Println("\t\t\t1.登陆聊天室")
		fmt.Println("\t\t\t2.注册用户")
		fmt.Println("\t\t\t3.退出系统")
		fmt.Println("\t\t\t请选择(1-3):")

		fmt.Scanf("%d\n", &key)
		switch key {
		case 1:
			fmt.Println("登陆聊天室")
			loop = false
		case 2:
			fmt.Println("注册用户")
			loop = false
		case 3:
			fmt.Println("退出系统")
			loop = false
		default:
			fmt.Println("你的输入有误，请重新输入")
		}
		if !loop {
			break
		}
	}

	if key == 1 {
		//用户登陆
		fmt.Println("请输入用户的ID")
		fmt.Scanf("%d\n", &userId)
		fmt.Println("请输入用户密码")
		fmt.Scanf("%s\n", &userPwd)
		login(userId, userPwd)
	}

	time.Sleep(20 * time.Second)
}
