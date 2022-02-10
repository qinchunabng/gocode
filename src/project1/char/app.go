package main

import "fmt"

//go默认使用的UTF-8编码
func main() {
	//go没有专门的字符类型，用byte存储字符
	var c1 byte = 'a'
	var c2 byte = '0'
	fmt.Println("c1=", c1, "c2=", c2)
	fmt.Printf("c1=%c,c2=%c\n", c1, c2)

	// var c3 byte = '北'
	var c3 int = '北'
	fmt.Printf("c3=%c,c3 code:%d\n", c3, c3)

	var c4 int = 22269
	fmt.Printf("c4=%c\n", c4)

	//字符可以当作整形运算
	var n1 = 10 + 'a'
	fmt.Printf("n1=%c", n1)
}
