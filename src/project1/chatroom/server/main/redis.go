package main

import (
	"github.com/garyburd/redigo/redis"
	"time"
)

var pool *redis.Pool

func initPool(address string, maxIdle int, maxActive int, idleTime time.Duration) {
	pool = &redis.Pool{
		MaxIdle:     maxIdle,   //最大空闲连接数
		MaxActive:   maxActive, //表示数据库的最大连接数，0表示没有限制
		IdleTimeout: idleTime,  //最大空闲时间
		Dial: func() (redis.Conn, error) {
			return redis.Dial("tcp", address)
		},
	}
}
