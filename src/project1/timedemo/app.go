package main

import (
	"fmt"
	"strconv"
	"time"
)

func main() {
	//获取当前时间
	now := time.Now()
	fmt.Printf("time type:%T, time val:%v", now, now)

	fmt.Printf("年=%v\n", now.Year())
	fmt.Printf("月=%v\n", now.Month())
	fmt.Printf("日=%v\n", now.Day())
	fmt.Printf("时=%v\n", now.Hour())
	fmt.Printf("分=%v\n", now.Minute())
	fmt.Printf("秒=%v\n", now.Second())

	//格式化日期和时间
	//方式一：
	str := fmt.Sprintf("当前时间： %d-%d-%d %d:%d:%d", now.Year(), now.Month(), now.Day(), now.Hour(), now.Minute(), now.Second())
	fmt.Println(str)
	//方式二：2006/01/02 15:04:05是固定的，但是可以自由组合
	fmt.Println(now.Format("2006/01/02 15:04:05"))
	//取月份
	fmt.Println(now.Format("01"))

	//每隔1秒打印一个数字，打印到100时就退出
	//i := 0
	//for {
	//	i++
	//	fmt.Println(i)
	//	//休眠
	//	time.Sleep(time.Millisecond * 100)
	//	if i == 100 {
	//		break
	//	}
	//}

	start := time.Now().Unix()
	test03()
	end := time.Now().Unix()
	fmt.Printf("test03()执行一共花了%v秒\n", end-start)
}

func test03() {
	str := ""
	for i := 0; i < 100000; i++ {
		str += "hello" + strconv.Itoa(i)
	}
}
