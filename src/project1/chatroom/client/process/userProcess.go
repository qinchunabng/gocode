package process

import (
	"encoding/json"
	"fmt"
	"net"
	"os"
	"project1/chatroom/common"
	"project1/chatroom/server/utils"
)

type UserProcess struct {
}

//写一个函数，完成登陆
func (this *UserProcess) Login(userId int, userPwd string) (err error) {
	fmt.Printf("userId=%v, userPwd=%v\n", userId, userPwd)
	//1.连接到服务器
	conn, err := net.Dial("tcp", "localhost:8889")
	if err != nil {
		fmt.Println("net.Dial err=", err)
		return err
	}
	defer conn.Close()
	//2.准备通过conn发送消息给服务器
	var mes common.Message
	mes.Type = common.LoginMesType
	//3.创建一个LoginMes结构体
	var loginMes common.LoginMes
	loginMes.UserId = userId
	loginMes.UserPwd = userPwd
	data, err := json.Marshal(loginMes)
	if err != nil {
		fmt.Println("json.Marshal err=", err)
		return err
	}
	//5.把消息体数据序列化复制给mes.Data
	mes.Data = string(data)
	//6.将mes序列化
	data, err = json.Marshal(mes)
	if err != nil {
		fmt.Println("json.Marshal err=", err)
		return err
	}
	//7.将数据发送给服务器
	tf := &utils.Transfer{
		Conn: conn,
	}
	err = tf.WritePkg(data)
	if err != nil {
		fmt.Println("登录WritePkg(data) err=", err)
	}
	mes, err = tf.ReadPkg()
	if err != nil {
		fmt.Println("登录readPkg(conn) err=", err)
		return
	}
	//将mes的Data反序列化为LoginResMes
	var loginResMes common.LoginResMes
	err = json.Unmarshal([]byte(mes.Data), &loginResMes)
	if err != nil {
		fmt.Println("json.Unmarshal err=", err)
		return
	}
	if loginResMes.Code == 200 {
		//初始化CurUser
		curUser.Conn = conn
		curUser.UserId = userId
		curUser.UserStatus = common.UserOnline

		//fmt.Println("登录成功")
		fmt.Println("当前在线用户列表如下：")
		for _, v := range loginResMes.UserIds {
			fmt.Println("用户ID：\t", v)
			user := &common.User{
				UserId:     v,
				UserStatus: common.UserOnline,
			}
			onlineUsers[v] = user
		}
		fmt.Println("\n")
		//使用协程保持客户端与服务器间的通信
		//接收并显示服务端发送的消息
		go processServerMes(conn)
		//循环显示菜单
		for {
			ShowMenu()
		}
	} else {
		fmt.Println(loginResMes.Error)
	}
	return
}

func (this *UserProcess) Register(userId int, userPwd string, userName string) (err error) {
	//1.连接到服务器
	conn, err := net.Dial("tcp", "localhost:8889")
	if err != nil {
		fmt.Println("net.Dial err=", err)
		return err
	}
	defer conn.Close()
	//2.准备通过conn发送消息给服务器
	var mes common.Message
	mes.Type = common.RegisterMesType
	//3.创建一个LoginMes结构体
	var registerMes common.RegisterMes
	registerMes.User.UserId = userId
	registerMes.User.UserName = userName
	registerMes.User.UserPwd = userPwd

	data, err := json.Marshal(registerMes)
	if err != nil {
		fmt.Println("json.Marshal err=", err)
		return err
	}

	//5.把消息体数据序列化复制给mes.Data
	mes.Data = string(data)
	//6.将mes序列化
	data, err = json.Marshal(mes)
	if err != nil {
		fmt.Println("json.Marshal err=", err)
		return err
	}
	//7.将数据发送给服务器
	tf := &utils.Transfer{
		Conn: conn,
	}
	err = tf.WritePkg(data)
	if err != nil {
		fmt.Println("注册WritePkg(data) fail. err=", err)
	}
	mes, err = tf.ReadPkg()
	if err != nil {
		fmt.Println("注册readPkg(conn) err=", err)
		return
	}
	//将mes的Data反序列化为LoginResMes
	var registerResMes common.RegisterResMes
	err = json.Unmarshal([]byte(mes.Data), &registerResMes)
	if err != nil {
		fmt.Println("json.Unmarshal err=", err)
		return
	}
	if registerResMes.Code == 200 {
		fmt.Println("注册成功，重新登录")
	} else {
		fmt.Println(registerResMes.Error)
	}
	//退出重新登录
	os.Exit(0)
	return
}
