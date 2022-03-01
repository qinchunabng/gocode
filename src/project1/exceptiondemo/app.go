package main

import (
	"errors"
	"fmt"
	"math/rand"
	"time"
)

//go中的异常处理，go中没有try...catch
//go中使用defer+panic+recover捕获处理异常
func main() {
	//test()
	//fmt.Println("执行下面的代码")

	test02()
}

func test() {
	defer func() {
		//recover是一个内置函数，它能捕获到异常
		err := recover()
		//如果不为空，则说明捕获到异常
		if err != nil {
			fmt.Println("err=", err)
		}
	}()
	num := 0
	res := 10 / num
	fmt.Println("res=", res)
}

func readConf(name string) (err error) {
	if name == "config.ini" {
		return nil
	} else {
		//返回一个自定义错误
		return errors.New("读取文件错误...")
	}
}

func test02() {
	err := readConf("config.ini")
	if err != nil {
		//如果读取文件发生错误，就输出这个错误，并终止这个程序
		panic(err)
	}
	fmt.Println("test02()继续执行。。。")
	rand.Seed(time.Now().Unix())
	for {
		n := rand.Intn(10) + 1
		fmt.Println("n=", n)
		if n == 10 {
			break
		}
		time.Sleep(time.Millisecond * 100)

	}

}
