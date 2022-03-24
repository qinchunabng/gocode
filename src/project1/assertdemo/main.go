package main

import "fmt"

func main() {
	//类型断言
	var t float32
	var x interface{}
	x = t
	y, ok := x.(int)
	if ok {
		fmt.Println("转换成功", y)
	} else {
		fmt.Println("转换失败")
	}

	var n1 float32 = 1.1
	var n2 float64 = 2.3
	var n3 int32 = 30
	var name string = "Tom"
	address := "北京"
	n4 := 300
	stu1 := Student{}
	stu2 := &Student{}
	TypeJudge(n1, n2, n3, name, address, n4, stu1, stu2, stu1, stu2)
}

type Student struct {
}

// TypeJudge 类型判断函数
func TypeJudge(items ...interface{}) {
	for i, item := range items {
		switch item.(type) { //type是关键字，固定写法，获取item的类型
		case bool:
			fmt.Printf("第%v个参数是bool类型，值是%v\n", i, item)
		case float32:
			fmt.Printf("第%v个参数是float32类型，值是%v\n", i, item)
		case float64:
			fmt.Printf("第%v个参数是float64类型，值是%v\n", i, item)
		case int, int32, int64:
			fmt.Printf("第%v个参数是整数类型，值是%v\n", i, item)
		case string:
			fmt.Printf("第%v个参数是string类型，值是%v\n", i, item)
		case Student:
			fmt.Printf("第%v个参数是Student类型，值是%v\n", i, item)
		case *Student:
			fmt.Printf("第%v个参数是*Student类型，值是%v\n", i, item)
		default:
			fmt.Printf("第%v个参数的类型不确定，值是%v\n", i, item)
		}
	}
}
