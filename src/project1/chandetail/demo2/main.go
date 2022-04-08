package main

import (
	"fmt"
	"strconv"
)

func main() {
	intChan := make(chan int, 10)
	for i := 0; i < 10; i++ {
		intChan <- i
	}
	stringChan := make(chan string, 5)
	for i := 0; i < 5; i++ {
		stringChan <- "hello" + strconv.Itoa(i)
	}

	//传统方法遍历管道时，常常会因为管道没有关闭导致deadlock的问题
	//在实际开发中可能无法确定何时关闭管道
	//可以使用select解决
	//label:
	for {
		select {
		case v := <-intChan:
			fmt.Printf("从intChan读取到值:%v\n", v)
		case v := <-stringChan:
			fmt.Printf("从stringChan读取到值:%v\n", v)
		default:
			fmt.Println("都取不到值，可以加自己的逻辑")
			return
			//break label
		}
	}
}
