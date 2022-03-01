package main

import "fmt"

func main() {
	num := new(int)
	*num = 100
	fmt.Printf("num的类型：%T,num的值：%v,num的地址:%v,指向的值:%v\n", num, num, &num, *num)
}
