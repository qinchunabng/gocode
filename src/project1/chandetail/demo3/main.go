package main

import (
	"fmt"
	"time"
)

func main() {
	go sayHello()
	go test()

	for i := 0; i < 10; i++ {
		fmt.Println("main() ok")
		time.Sleep(time.Second)
	}
}

func sayHello() {
	for i := 0; i < 10; i++ {
		time.Sleep(time.Second)
		fmt.Println("hello world")
	}
}
func test() {
	//使用异常捕获防止异常导致整个程序崩溃
	defer func() {
		//捕获panic
		if err := recover(); err != nil {
			fmt.Println("test()发生错误,err=", err)
		}
	}()
	var myMap map[int]string
	myMap[0] = "golang"
}
