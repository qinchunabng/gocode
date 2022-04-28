package main

import (
	"encoding/binary"
	"encoding/json"
	"fmt"
	"io"
	"net"
	"project1/chatroom/common"
)

func main() {
	fmt.Println("服务器在8889端口监听...")
	listen, err := net.Listen("tcp", "0.0.0.0:8889")
	if err != nil {
		fmt.Println("net.Listen err=", err)
		return
	}
	defer listen.Close()

	//等待客户端连接
	for {
		fmt.Println("等待客户端连接服务器...")
		conn, err := listen.Accept()
		if err != nil {
			fmt.Println("listen.Accept err=", err)
		}
		//一旦连接成功，则启动一个协程和客户端保持通讯
		go process(conn)
	}
}

//处理客户端通信
func process(conn net.Conn) {
	defer conn.Close()
	//读取客户端发送的消息
	for {
		mes, err := readPkg(conn)
		if err != nil {
			if err == io.EOF {
				fmt.Println("客户端退出，服务端也退出...")
				return
			}
			fmt.Println("readPkg err=", err)
			return
		}
		//fmt.Println("mes=", mes)
		err = serverProcessMes(conn, &mes)
		if err != nil {
			return
		}
	}
}

//读取客户端发送过来的消息
func readPkg(conn net.Conn) (mes common.Message, err error) {
	buf := make([]byte, 8096)
	fmt.Println("读取客户端发送的数据...")
	_, err = conn.Read(buf[:4])
	if err != nil {
		fmt.Println("read pkg header error, err=", err)
		return
	}

	//根据buf[:4]转换为一个uint32
	var pkgLen uint32
	pkgLen = binary.BigEndian.Uint32(buf[:4])

	//根据pkgLen读取消息内容
	n, err := conn.Read(buf[:pkgLen])
	if uint32(n) != pkgLen && err != nil {
		fmt.Println("read pkg error. err=", err)
		//err = errors.New("read pkg error. msg: " + err.Error())
		return
	}

	//反序列化
	err = json.Unmarshal(buf[:pkgLen], &mes)
	if err != nil {
		fmt.Println("pkg unmarshal error. err=", err)
		//err = errors.New("pkg unsearialize error. msg: " + err.Error())
		return
	}
	return
}

//服务端消息处理函数
func serverProcessMes(conn net.Conn, mes *common.Message) (err error) {
	switch mes.Type {
	case common.LoginMesType:
		err = serverProcessLogin(conn, mes)
		//处理登录的逻辑
	case common.RegisterMesType:
	//处理注册
	default:
		fmt.Println("消息类型不存在，无法处理...")
	}
	return
}

//处理登录的业务逻辑
func serverProcessLogin(conn net.Conn, mes *common.Message) (err error) {
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
	//校验账号密码
	if loginMes.UserId == 100 && loginMes.UserPwd == "123456" {
		loginResMes.Code = 200
	} else {
		loginResMes.Code = 500
		loginResMes.Error = "该用户不存在，请注册再使用..."
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
	err = writePkg(conn, data)
	return
}

func writePkg(conn net.Conn, data []byte) (err error) {
	//7.1.先把data的长度发送给服务器
	//先获取data的长度->转换为一个标识长度的byte切片
	var pkgLen uint32
	pkgLen = uint32(len(data))
	var bytes [4]byte
	binary.BigEndian.PutUint32(bytes[0:4], pkgLen)
	//发送长度
	n, err := conn.Write(bytes[:])
	if n != 4 || err != nil {
		fmt.Println("conn.Write(bytes) fail", err)
		return
	}

	//发送data本身
	n, err = conn.Write(data)
	if uint32(n) != pkgLen || err != nil {
		fmt.Println("conn.Write(bytes) fail", err)
		return
	}
	return
}
