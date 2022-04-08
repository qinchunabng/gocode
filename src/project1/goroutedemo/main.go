package main

import (
	"fmt"
	"strconv"
	"sync"
	"time"
)

var (
	myMap = make(map[int]int)
	//声明一个互斥锁
	lock sync.Mutex
)

func main() {
	//启动一个协程
	//协程运行规则，主线程结束，协程没有执行完也会结束
	//如果主线程没有结束，协程执行完毕也会结束
	//go test()
	//for i := 0; i < 10; i++ {
	//	fmt.Println("main() : Hello World! " + strconv.Itoa(i+1))
	//	time.Sleep(time.Second)
	//}

	//开启200个协程求1到200的阶乘
	for i := 1; i <= 20; i++ {
		go factorial(i)
	}

	time.Sleep(time.Second * 5)

	lock.Lock()
	for k, v := range myMap {
		fmt.Printf("map[%d]=%d\n", k, v)
	}
	lock.Unlock()
}

//求阶乘，n!
func factorial(n int) {
	res := 1
	for i := 1; i <= n; i++ {
		res *= i
	}
	//加锁
	lock.Lock()
	myMap[n] = res
	//解锁
	lock.Unlock()
}

func test() {
	for i := 0; i < 10; i++ {
		fmt.Println("test() : Hello World! " + strconv.Itoa(i+1))
		time.Sleep(time.Second)
	}
}
