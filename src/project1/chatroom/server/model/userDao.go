package model

import (
	"encoding/json"
	"fmt"
	"github.com/garyburd/redigo/redis"
	"project1/chatroom/common"
)

type UserDao struct {
	pool *redis.Pool
}

//声明一个全局变量
var (
	MyUserDao *UserDao
)

// NewUserDao 使用工程模式创建一个userDao的实例
func NewUserDao(pool *redis.Pool) (userDao *UserDao) {
	userDao = &UserDao{
		pool,
	}
	return
}

//通过id查询user
func (this *UserDao) getUserById(conn redis.Conn, id int) (user *common.User, err error) {
	res, err := redis.String(conn.Do("HGet", "users", id))
	if err != nil {
		if err == redis.ErrNil {
			//表示redis不存在对应的数据
			err = ERROR_USER_NOTEXISTS
		}
		return
	}

	user = &common.User{}
	//反序列化
	err = json.Unmarshal([]byte(res), user)
	if err != nil {
		fmt.Println("json.Unmarshal err=", err)
		return
	}
	return
}

// Login 完成登陆的校验
func (this *UserDao) Login(userId int, userPwd string) (user *common.User, err error) {
	conn := this.pool.Get()
	defer conn.Close()
	user, err = this.getUserById(conn, userId)
	if err != nil {
		return
	}

	//判断用户密码是否正确
	if user.UserPwd != userPwd {
		err = ERROR_USER_PWD
		return
	}
	return
}

// Register 注册用户
func (this *UserDao) Register(user *common.User) (err error) {
	conn := this.pool.Get()
	defer conn.Close()
	u, err := this.getUserById(conn, user.UserId)
	if err != nil && err != ERROR_USER_NOTEXISTS {
		return
	} else if u != nil {
		//用户已存在
		err = ERROR_USER_EXISTS
		return
	}
	//序列化
	data, err := json.Marshal(user)
	if err != nil {
		fmt.Println("json.Marshal err=", err)
		return
	}
	//入库
	_, err = conn.Do("HSet", "users", user.UserId, string(data))
	if err != nil {
		fmt.Println("保存注册用户错误,err=", err)
		return
	}
	return
}
