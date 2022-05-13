package model

import (
	"net"
	"project1/chatroom/common"
)

type CurUser struct {
	Conn net.Conn
	common.User
}
