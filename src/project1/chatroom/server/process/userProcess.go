package process

import (
	"encoding/json"
	"fmt"
	"net"
	"project1/chatroom/common"
	"project1/chatroom/server/model"
	"project1/chatroom/server/utils"
)

type UserProcess struct {
	Conn   net.Conn
	UserId int
}

// ServerProcessLogin 处理登录的业务逻辑
func (this *UserProcess) ServerProcessLogin(mes *common.Message) (err error) {
	//1.先从mes中取出mes.Data，并直接反序列化为LoginMes
	var loginMes common.LoginMes
	err = json.Unmarshal([]byte(mes.Data), &loginMes)
	if err != nil {
		fmt.Println("json.Unmarshal fail err=", err)
		return
	}
	//先声明一个resMes
	var resMes common.Message
	resMes.Type = common.LoginResMesType

	var loginResMes common.LoginResMes
	//使用UserDao到redis中验证
	user, err := model.MyUserDao.Login(loginMes.UserId, loginMes.UserPwd)
	if err != nil {
		loginResMes.Code = 500
		if err == model.ERROR_USER_NOTEXISTS {
			loginResMes.Error = err.Error()
		} else if err == model.ERROR_USER_PWD {
			loginResMes.Error = err.Error()
		} else {
			loginResMes.Error = "服务器内部错误..."
		}

	} else {
		loginResMes.Code = 200
		fmt.Println(user, "登陆成功")
		//用户登陆成功，添加userProcess
		this.UserId = loginMes.UserId
		userMgr.AddOnlineUser(this)
		//通知其他用户我上线了
		this.NotifyOtherOnlineUser(this.UserId)

		for id, _ := range userMgr.onlineUsers {
			loginResMes.UserIds = append(loginResMes.UserIds, id)
		}
	}

	//序列化
	data, err := json.Marshal(loginResMes)
	if err != nil {
		fmt.Println("json.Marshal fail", err)
		return
	}
	resMes.Data = string(data)
	//对resMes进行序列化
	data, err = json.Marshal(resMes)
	if err != nil {
		fmt.Println("json.Marshal fail", err)
		return
	}
	//发送数据
	transfer := &utils.Transfer{
		Conn: this.Conn,
	}
	err = transfer.WritePkg(data)
	return
}

// ServerProcessRegister 处理注册的业务逻辑
func (this *UserProcess) ServerProcessRegister(mes *common.Message) (err error) {
	//1.先从mes中取出mes.Data，并直接反序列化为LoginMes
	var registerMes common.RegisterMes
	err = json.Unmarshal([]byte(mes.Data), &registerMes)
	if err != nil {
		fmt.Println("json.Unmarshal fail err=", err)
		return
	}
	//先声明一个resMes
	var resMes common.Message
	resMes.Type = common.RegisterResMesType

	var registerResMes common.RegisterResMes
	//使用UserDao到redis中验证
	err = model.MyUserDao.Register(&registerMes.User)
	if err != nil {
		//注册失败
		registerResMes.Code = 500
		registerResMes.Error = err.Error()
	} else {
		registerResMes.Code = 200
	}

	//发送响应消息
	//序列化
	data, err := json.Marshal(registerResMes)
	if err != nil {
		fmt.Println("json.Marshal fail", err)
		return
	}
	resMes.Data = string(data)
	//对resMes进行序列化
	data, err = json.Marshal(resMes)
	if err != nil {
		fmt.Println("json.Marshal fail", err)
		return
	}
	//发送数据
	transfer := &utils.Transfer{
		Conn: this.Conn,
	}
	err = transfer.WritePkg(data)
	return
}

// NotifyOtherOnlineUser 通知其他在线用户
func (this *UserProcess) NotifyOtherOnlineUser(userId int) {
	for id, up := range userMgr.onlineUsers {
		if id == userId {
			continue
		}
		//通知上线
		up.NotifyMeOnline(userId)
	}
}

// NotifyMeOnline 上线通知
func (this *UserProcess) NotifyMeOnline(userId int) {
	var mes common.Message
	mes.Type = common.NotifyUserStatusMesType

	var notificationMes common.NotifyUserStatusMes
	notificationMes.UserId = userId
	notificationMes.Status = common.UserOnline

	//序列化数据
	data, err := json.Marshal(notificationMes)
	if err != nil {
		fmt.Println("json.Marshal notificationMes err=", err)
		return
	}
	mes.Data = string(data)

	data, err = json.Marshal(mes)
	if err != nil {
		fmt.Println("json.Marshal mes err=", err)
		return
	}
	transfer := &utils.Transfer{
		Conn: this.Conn,
	}
	err = transfer.WritePkg(data)
	if err != nil {
		fmt.Println("NotifyMeOnline err = ", err)
	}
}
