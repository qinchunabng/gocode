package process

import (
	"encoding/json"
	"fmt"
	"net"
	"project1/chatroom/common"
	"project1/chatroom/server/utils"
)

type SmsProcess struct {
}

// SendGroupMes 转发消息
func (this *SmsProcess) SendGroupMes(mes *common.Message) {
	var smsMes common.SmsMes
	err := json.Unmarshal([]byte(mes.Data), &smsMes)
	if err != nil {
		fmt.Println("SendGroupMes json.Unmarshal err=", err)
		return
	}
	data, err := json.Marshal(mes)
	if err != nil {
		fmt.Println("SendGroupMes json.Unmarshal err=", err)
		return
	}
	for id, up := range userMgr.onlineUsers {
		if id == smsMes.UserId {
			continue
		}
		this.SendMesToEachOnlineUser(data, up.Conn)
	}
}

// SendMesToEachOnlineUser 发送消息给对应的各个用户
func (this *SmsProcess) SendMesToEachOnlineUser(data []byte, conn net.Conn) {
	tf := &utils.Transfer{
		Conn: conn,
	}
	err := tf.WritePkg(data)
	if err != nil {
		fmt.Println("转发消息失败err=", err)
	}
}
