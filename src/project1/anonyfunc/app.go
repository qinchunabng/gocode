package main

import "fmt"

var (
	//全局匿名函数
	func1 = func(n1 int, n2 int) int {
		return n1 * n2
	}
)

func main() {
	//匿名函数第一种用法
	res1 := func(n1 int, n2 int) int {
		return n1 + n2
	}(10, 20)
	fmt.Printf("res1=%v\n", res1)
	//匿名函数第二种用法
	a := func(n1 int, n2 int) int {
		return n1 - n2
	}
	fmt.Printf("a=%T\n", a)
	res2 := a(10, 20)
	fmt.Printf("res2=%v\n", res2)

	res3 := func1(10, 20)
	fmt.Printf("res3=%v\n", res3)
}
