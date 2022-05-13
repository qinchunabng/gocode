package process

import (
	"fmt"
	"project1/chatroom/client/model"
	"project1/chatroom/common"
)

//客户端存储在线用户
var onlineUsers map[int]*common.User = make(map[int]*common.User, 10)

//在登陆成功后完成初始化
var curUser model.CurUser

//更新用户状态
func updateUserStatus(notifyUserStatusMes *common.NotifyUserStatusMes) {
	user, ok := onlineUsers[notifyUserStatusMes.UserId]
	if !ok {
		user = &common.User{
			UserId:     notifyUserStatusMes.UserId,
			UserStatus: notifyUserStatusMes.Status,
		}
	}
	user.UserStatus = notifyUserStatusMes.Status
	onlineUsers[notifyUserStatusMes.UserId] = user

	outputOnlineUser()
}

//显示当前在线用户
func outputOnlineUser() {
	fmt.Println("当前在线用户列表:")
	for id, _ := range onlineUsers {
		fmt.Println("用户id:\t", id)
	}
}
