package main

import "fmt"

func main() {
	for i := 1; i <= 10; i++ {
		fmt.Println("Hello World")
	}

	//for第二种写法
	j := 1
	for j <= 10 {
		fmt.Println("Hello World~")
		j++ //循环变量的迭代
	}

	//死循环，相当于for;;{},通常需要配合一个break语句
	k := 1
	for {
		//循环语句
		if k <= 10 {
			fmt.Println("ok")
		} else {
			break
		}
		k++
	}

	//字符串遍历方式一
	//如果字符串含有中文，传统的遍历方式会出现乱码，传统的遍历的时候是按照字节遍历的
	//而UTF-8编码一个中文占三个字节，解决方法是转为切片
	var str string = "hello,world!中国"
	str2 := []rune(str)
	for i := 1; i < len(str2); i++ {
		fmt.Printf("%c\n", str2[i])
	}

	//字符串遍历方式二
	//按照字符的方式遍历
	str = "abc~ok武汉"
	for index, val := range str {
		fmt.Printf("index=%d, val=%c\n", index, val)
	}
}
