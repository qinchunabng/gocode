package main

import (
	"fmt"
	"unsafe"
)

var (
	n3    = 300
	n4    = 900
	name2 = "mary"
)

func main() {
	// fmt.Println("hello,world!")
	var i int
	fmt.Println("i=", i)

	var num = 10.11
	fmt.Println("num=", num)

	name := "tom"
	fmt.Println("name=", name)

	// var n1, n2, n3 int
	// fmt.Println("n1=", n1, "n2=", n2, "n3=", n3)

	// var n1, n2, n3 = 100, "tom", 888
	// fmt.Println("n1=", n1, "n2=", n2, "n3=", n3)

	n1, n2, n3 := 100, "tom~", 888
	fmt.Println("n1=", n1, "n2=", n2, "n3=", n3)

	fmt.Println("n3=", n3, "n4=", n4, "name2=", name2)

	//查看变量的类型和占用的字节数
	var t1 int = 100
	fmt.Printf("n1类型: %T，占用的字节数：%d", t1, unsafe.Sizeof(t1))
}
