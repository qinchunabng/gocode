package utils

import (
	"encoding/binary"
	"encoding/json"
	"fmt"
	"net"
	"project1/chatroom/common"
)

type Transfer struct {
	Conn net.Conn
	Buf  [8096]byte
}

//读取客户端发送过来的消息
func (this *Transfer) ReadPkg() (mes common.Message, err error) {
	fmt.Println("读取客户端发送的数据...")
	_, err = this.Conn.Read(this.Buf[:4])
	if err != nil {
		fmt.Println("read pkg header error, err=", err)
		return
	}

	//根据buf[:4]转换为一个uint32
	var pkgLen uint32
	pkgLen = binary.BigEndian.Uint32(this.Buf[:4])

	//根据pkgLen读取消息内容
	n, err := this.Conn.Read(this.Buf[:pkgLen])
	if uint32(n) != pkgLen && err != nil {
		fmt.Println("read pkg error. err=", err)
		//err = errors.New("read pkg error. msg: " + err.Error())
		return
	}

	//反序列化
	err = json.Unmarshal(this.Buf[:pkgLen], &mes)
	if err != nil {
		fmt.Println("pkg unmarshal error. err=", err)
		//err = errors.New("pkg unsearialize error. msg: " + err.Error())
		return
	}
	return
}

func (this *Transfer) WritePkg(data []byte) (err error) {
	//7.1.先把data的长度发送给服务器
	//先获取data的长度->转换为一个标识长度的byte切片
	var pkgLen uint32
	pkgLen = uint32(len(data))
	binary.BigEndian.PutUint32(this.Buf[0:4], pkgLen)
	//发送长度
	n, err := this.Conn.Write(this.Buf[:4])
	if n != 4 || err != nil {
		fmt.Println("conn.Write(bytes) fail", err)
		return
	}

	//发送data本身
	n, err = this.Conn.Write(data)
	if uint32(n) != pkgLen || err != nil {
		fmt.Println("conn.Write(bytes) fail", err)
		return
	}
	return
}
