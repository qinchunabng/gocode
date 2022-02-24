package utils

import (
	"fmt"
	"project1/packagedemo/test" //循环引用会报错
)

var Test = test.Test

//将函数首字符大写标识该函数可以外部使用
func Calc(n1 float64, n2 float64, operate byte) float64 {
	var res float64
	switch operate {
	case '+':
		res = n1 + n2
	case '-':
		res = n1 - n2
	case '*':
		res = n1 * n2
	case '/':
		res = n1 / n2
	default:
		fmt.Println("操作符号错误...")
	}
	return res
}
