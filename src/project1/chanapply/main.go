package main

import (
	"fmt"
	"time"
)

func main() {
	intChan := make(chan int, 50)
	exitChan := make(chan bool, 1)

	//开启一个写协程，往管道里面读数据
	go writeData(intChan)
	//开启一个读协程，读取管道里面的数据
	go readData(intChan, exitChan)

	for {
		v, ok := <-exitChan
		fmt.Printf("v=%v,ok=%v\n", v, ok)
		if !ok {
			break
		}
	}
	//time.Sleep(time.Second * 5)
}

func writeData(intChan chan int) {
	for i := 0; i < 50; i++ {
		//向管道放入数据
		intChan <- i
		fmt.Printf("写入数据%v\n", i)
		time.Sleep(time.Second)
	}
	close(intChan)
}

func readData(intChat chan int, exitChan chan bool) {
	for {
		v, ok := <-intChat
		if !ok {
			break
		}
		fmt.Printf("readData读到数据=%v\n", v)
		time.Sleep(time.Second)
	}
	//读取数据完成
	exitChan <- true
	close(exitChan)
}
