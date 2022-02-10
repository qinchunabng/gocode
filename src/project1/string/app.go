package main

import "fmt"

func main() {
	var s string = "hello world"
	fmt.Println(s)
	//go中的字符串是不可变的
	//反引号的使用
	str := `\n`
	fmt.Println("str=", str)

	//字符串拼接
	str += "haha"
	fmt.Println("str=", str)

	//字符串多行拼接，加号保留在上一行
	var str1 = "hello" +
		" world" +
		" hello"
	fmt.Println("str1=" + str1)

	//基本类型初始值
	var i int
	var b1 byte
	var f float32
	var f1 float64
	var b bool
	var str2 string
	//这里的%v按照变量的原始值输出
	fmt.Printf("i=%d,b1=%v,f=%v,f1=%v,b=%v,str2=%v", i, b1, f, f1, b, str2)
}
