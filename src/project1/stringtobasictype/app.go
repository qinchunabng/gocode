package main

import (
	"fmt"
	"strconv"
)

//string转基本类型
func main() {
	var str string = "true"
	var b bool
	//1.strconv.ParseBool会返回两个值（value bool, err error）
	//2.因为只想获取到value bool,所以使用_忽略err error
	b, _ = strconv.ParseBool(str)
	fmt.Printf("b type %T b=%v\n", b, b)

	var str2 string = "1234567890"
	var n1 int64
	n1, _ = strconv.ParseInt(str2, 10, 64)
	fmt.Printf("b type %T b=%v\n", n1, n1)

	//string转基本数据类型，如果不是有效数据，会转换为对应数据的类型的初始值
	var str3 string = "hello"
	var n2 int64
	n2, _ = strconv.ParseInt(str3, 10, 64)
	fmt.Printf("b type %T b=%v\n", n2, n2)
}
