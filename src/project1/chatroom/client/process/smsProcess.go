package process

import (
	"encoding/json"
	"fmt"
	"project1/chatroom/common"
	"project1/chatroom/server/utils"
)

type SmsProcess struct {
}

func (this *SmsProcess) SendGroupMes(content string) (err error) {
	//创建Mes
	var mes common.Message
	mes.Type = common.SmsMesType

	//创建SmsMes实例
	var smsMes common.SmsMes
	smsMes.Content = content
	smsMes.User = curUser.User

	//序列化smsMes
	data, err := json.Marshal(smsMes)
	if err != nil {
		fmt.Println("SendGroupMes json.Marshal fail = ", err)
	}
	mes.Data = string(data)

	//将mes序列化
	data, err = json.Marshal(mes)
	if err != nil {
		fmt.Println("SendGroupMes json.Marshal fail=", err)
	}

	//发送给服务器
	transfer := &utils.Transfer{
		Conn: curUser.Conn,
	}
	err = transfer.WritePkg(data)
	if err != nil {
		fmt.Println("SendGroupMes err=", err)
		return
	}
	return
}
