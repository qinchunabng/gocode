package process

import (
	"encoding/json"
	"fmt"
	"project1/chatroom/common"
)

func outputGroupMes(mes *common.Message) {
	var smsMes *common.SmsMes
	err := json.Unmarshal([]byte(mes.Data), &smsMes)
	if err != nil {
		fmt.Println("outputGroupMes json.Unmarshal err=", err)
		return
	}

	//显示信息
	fmt.Printf("用户ID: %v 对大家说: %v\n", smsMes.UserId, smsMes.Content)
}
