package main

import "fmt"

func main() {
	res := sum(10, 20)
	fmt.Println("res=", res)
}

//求和
func sum(n1 int, n2 int) int {
	//当执行到defer时，会将defer后的语句压入到栈中，当函数执行完之后，按照先入后出的方式执行
	defer fmt.Println("ok1 n1=", n1)
	defer fmt.Println("ok2 n2=", n2)
	res := n1 + n2
	fmt.Println("ok3 res=", res)
	return res
}
