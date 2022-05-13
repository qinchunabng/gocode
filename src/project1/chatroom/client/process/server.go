package process

import (
	"encoding/json"
	"fmt"
	"net"
	"os"
	"project1/chatroom/common"
	"project1/chatroom/server/utils"
)

// ShowMenu 显示登陆后的成功的界面
func ShowMenu() {
	fmt.Println("---------恭喜xxxx登陆成功---------")
	fmt.Println("---------1. 显示在线用户列表-----------")
	fmt.Println("---------2. 发送消息-----------")
	fmt.Println("---------3. 信息列表-----------")
	fmt.Println("---------4. 退出系统-----------")
	fmt.Println("请选择(1-4):")
	var key int
	var content string
	smsProcess := &SmsProcess{}
	fmt.Scanf("%d\n", &key)
	switch key {
	case 1:
		//fmt.Println("显示在线用户列表.")
		outputOnlineUser()
	case 2:
		fmt.Println("请输入你想对大家说的话：")
		fmt.Scanf("%s\n", &content)
		smsProcess.SendGroupMes(content)
	case 3:
		fmt.Println("信息列表.")
	case 4:
		fmt.Println("你选择了退出系统.")
		os.Exit(0)
	default:
		fmt.Println("你输入的选项不对.")
	}
}

//客户端和服务器保持通讯
func processServerMes(conn net.Conn) {
	//创建一个transfer实例不停读取服务器发送的消息
	tf := &utils.Transfer{
		Conn: conn,
	}
	for {
		fmt.Printf("客户端正%v在读取服务器发送的消息\n", conn.LocalAddr().String())
		mes, err := tf.ReadPkg()
		if err != nil {
			fmt.Println("tf.ReadPkg err=", err)
		}
		//如果读取到消息
		//fmt.Printf("mes=%v\n", mes)
		switch mes.Type {
		case common.NotifyUserStatusMesType:
			//有人上线了
			var notifyUserStatusMes common.NotifyUserStatusMes
			json.Unmarshal([]byte(mes.Data), &notifyUserStatusMes)
			//更新用户状态
			updateUserStatus(&notifyUserStatusMes)
		case common.SmsMesType:
			outputGroupMes(&mes)
		default:
			fmt.Println("服务器返回了未知消息类型")
		}
	}
}
