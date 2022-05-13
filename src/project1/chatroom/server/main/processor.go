package main

import (
	"fmt"
	"io"
	"net"
	"project1/chatroom/common"
	process2 "project1/chatroom/server/process"
	"project1/chatroom/server/utils"
)

type Processor struct {
	Conn net.Conn
}

//服务端消息处理函数
func (this *Processor) serverProcessMes(mes *common.Message) (err error) {
	fmt.Println("mes=", mes)
	switch mes.Type {
	case common.LoginMesType:
		//处理登录的逻辑
		up := &process2.UserProcess{
			Conn: this.Conn,
		}
		err = up.ServerProcessLogin(mes)
	case common.RegisterMesType:
		//处理注册的逻辑
		up := &process2.UserProcess{
			Conn: this.Conn,
		}
		err = up.ServerProcessRegister(mes)
	case common.SmsMesType:
		//创建一个SmsProcess实例完成转发群聊消息
		sp := &process2.SmsProcess{}
		sp.SendGroupMes(mes)
	default:
		fmt.Println("消息类型不存在，无法处理...")
	}
	return
}

//处理客户端通信
func (this *Processor) Process() (err error) {
	//读取客户端发送的消息
	for {
		transfer := &utils.Transfer{
			Conn: this.Conn,
		}
		mes, err := transfer.ReadPkg()
		if err != nil {
			if err == io.EOF {
				fmt.Println("客户端退出，服务端也退出...")
				return err
			}
			fmt.Println("readPkg err=", err)
			return err
		}
		//fmt.Println("mes=", mes)
		err = this.serverProcessMes(&mes)
		if err != nil {
			return err
		}
	}
}
