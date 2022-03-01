package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	//go默认按照UTF-8格式编码，中文占三个字节，len返回的是字节数
	str := "hello中"
	fmt.Println("str len:", len(str))

	str2 := "hello北京"
	//遍历含中文的字符串
	r := []rune(str2)
	for i := 0; i < len(r); i++ {
		fmt.Printf("字符=%c\n", r[i])
	}

	//字符串转整数
	n, err := strconv.Atoi("hell")
	if err != nil {
		fmt.Println("转换错误", err)
	} else {
		fmt.Printf("转换的结果是:%v\n", n)
	}

	//4）整数转字符串
	str = strconv.Itoa(12345)
	fmt.Printf("str type:%T,str:%v\n", str, str)

	//5）字符串转字节数组
	var bytes = []byte("hello go")
	fmt.Printf("bytes:%v\n", bytes)

	//6)byte数组转字符串
	str = string([]byte{97, 98, 99})
	fmt.Printf("str=%v\n", str)

	//7)10进制转2，8，16进制，返回对应的字符串
	str = strconv.FormatInt(123, 2)
	fmt.Printf("123对应的二进制是:%v\n", str)
	str = strconv.FormatInt(123, 16)
	fmt.Printf("123对应的十六进制是:%v\n", str)

	//8)查找字符串
	var b = strings.Contains("seafood", "foo")
	fmt.Printf("b=%v\n", b)

	//9)统计包含子字符的个数
	var count = strings.Count("cheese", "e")
	fmt.Printf("count=%v\n", count)

	//10)字符比较
	//==区分大小写
	b = "Abc" == "abc"
	fmt.Printf("b=%v\n", b)
	//strings.EqualFold不分区大小写
	b = strings.EqualFold("Abc", "abc")
	fmt.Printf("b=%v\n", b)

	//11)返回子串第一次出现的下标位置,没找到返回-1
	index := strings.Index("NTL_abc", "abc")
	fmt.Printf("index=%v\n", index)

	//12)返回字串最后一次出现的未知
	index = strings.LastIndex("golang go", "go")
	fmt.Printf("index=%v\n", index)

	//13)替换字符串，最后一个参数为替换几个，为-1替换所有
	//被替换的字符串并没有发生变化
	str = strings.Replace("golang go", "go", "java", -1)
	fmt.Printf("str=%v\n", str)

	//14)字符串拆分
	strArr := strings.Split("hello,world,ok", ",")
	for i := 0; i < len(strArr); i++ {
		fmt.Printf("strArr[%v]=%v\n", i, strArr[i])
	}
	fmt.Printf("strArr=%v\n", strArr)

	//15)大小写转换
	fmt.Printf("转小写：%v\n", strings.ToLower("Go"))
	fmt.Printf("转大写: %v\n", strings.ToUpper("Go"))

	//16)去掉字符串左右的空格
	str = strings.TrimSpace(" hello ")
	fmt.Printf("str=%q\n", str)

	//17)去掉左右指定的内容的字符串
	str = strings.Trim("!hello!", "!")
	fmt.Printf("str=%q\n", str)

	//18)判断字符串是否以指定字符串开头或结尾
	b = strings.HasPrefix("http://www.baidu.com", "http")
	fmt.Printf("b=%v\n", b)
	b = strings.HasSuffix("http://www.baidu.com", ".com")
	fmt.Printf("b=%b\n", b)
}
