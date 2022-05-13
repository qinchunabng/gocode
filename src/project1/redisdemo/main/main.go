package main

import (
	"fmt"
	"github.com/garyburd/redigo/redis"
)

func main() {
	//通过go向redis写入数据和读取数据
	//1.连接到redis
	//导入redis包前，需要先安装redis包，然后在使用go mod tidy命令将依赖添加到module
	conn, err := redis.Dial("tcp", "127.0.0.1:6379")
	if err != nil {
		fmt.Println("redis.Dail err=", err)
		return
	}
	defer conn.Close()

	//2.通过go向redis写入数据
	_, err = conn.Do("set", "name", "tom")
	if err != nil {
		fmt.Println("set err=", err)
		return
	}

	val, err := redis.String(conn.Do("get", "name"))
	if err != nil {
		fmt.Println("get err=", err)
		return
	}
	fmt.Printf("val=%v\n", val)
	fmt.Println("conn succ...", conn)
}
