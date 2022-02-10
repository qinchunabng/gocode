package main

import "fmt"

func main() {
	//基本类型的数据转换
	var i int32 = 100
	//将i转为float
	var n1 float32 = float32(i)
	var n2 int64 = int64(i)
	fmt.Println("n1=", n1, "n2=", n2)

	//基本数据类型转string

}
