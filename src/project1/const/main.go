package main

import "fmt"

func main() {
	//1.常量声明时必须赋值
	//2.常量定义好不能修改
	//3.常量只能修饰bool,int,bool,string基本类型
	const tax = 1
	fmt.Println("tax=", tax)
	//
	const c2 = tax * 2
	fmt.Printf("c2=%v\n", c2)

	//这种写法不可以，常量必须是在编译时就确定的常量
	//num2 := 10
	//const c1 = num2 * 2
	//fmt.Printf("c1=%v\n", c1)

	//常量的简洁写法一
	const (
		a = 1
		b = 2
	)
	fmt.Println(a, b)

	//常量的简洁写法二
	//表示将c赋值为0
	//后面的在前面的基础上+1
	const (
		c = iota
		d
		e
	)
	fmt.Println(c, d, e)

	//在go中仍然通过首字母的大小写控制访问
}
