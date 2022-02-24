package main

import (
	"fmt"
	"strings"
)

func main() {
	f := AddUpper()
	fmt.Println(f(1))
	fmt.Println(f(2))

	f1 := makeSuffix(".jpg")
	fmt.Printf("文件处理后=%v\n", f1("bird.jpg"))
}

//累加器
func AddUpper() func(int) int {
	var n = 10
	var str = "hello"
	return func(x int) int {
		n = n + x
		str += fmt.Sprint(x)
		fmt.Printf("str=%v\n", str)
		return n
	}
}

//判断文件名是否有指定后缀，如果有则直接返回
//没有则拼接上后缀
func makeSuffix(suffix string) func(string) string {
	return func(name string) string {
		if strings.HasSuffix(name, suffix) {
			return name
		} else {
			return name + suffix
		}
	}
}
