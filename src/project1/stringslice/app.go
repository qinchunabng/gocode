package main

import "fmt"

func main() {
	str := "hello,go"
	//修改字符串内容
	arr1 := []byte(str)
	arr1[0] = 'z'
	str = string(arr1)
	fmt.Println("str=", str)

	//修改包含汉字的字符串
	str = "hello,中国"
	arr2 := []rune(str)
	arr2[6] = '美'
	str = string(arr2)
	fmt.Println("str=", str)
}
