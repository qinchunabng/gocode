package main

import (
	"fmt"
	"reflect"
)

func main() {
	var s Monster = Monster{
		Name:  "黄鼠狼精",
		Age:   400,
		Score: 30.8,
	}

	TestStruct(s)
}

func TestStruct(a interface{}) {
	typ := reflect.TypeOf(a)
	val := reflect.ValueOf(a)
	kd := val.Kind()
	//判断kind是否是结构体
	if kd != reflect.Struct {
		fmt.Println("expect struct")
		return
	}

	//获取该结构体有几个字段
	num := val.NumField()
	fmt.Printf("struct has %d fields\n", num)
	//遍历结构体的所有字段
	for i := 0; i < num; i++ {
		fmt.Printf("Field %d 值为=%v\n", i, val.Field(i))
		//获取struct的标签，需要用reflect.Type来获取tag标签的值
		tagVal := typ.Field(i).Tag.Get("json")
		if tagVal != "" {
			fmt.Printf("Field %d: tag为=%v\n", i, tagVal)
		}
	}
	numOfMethod := val.NumMethod()
	fmt.Printf("struct has %d methods\n", numOfMethod)
	val.Method(1).Call(nil)

	//调用结构体的第一个方法
	//方法的排序按照函数名称的ASCII码排序
	var params []reflect.Value
	params = append(params, reflect.ValueOf(10))
	params = append(params, reflect.ValueOf(20))
	res := val.Method(0).Call(params)
	fmt.Println("res=", res[0].Int())
}

type Monster struct {
	Name  string `json:"name"`
	Age   int    `json:"age"`
	Score float32
	Sex   string
}

func (s Monster) Print() {
	fmt.Println("---start---")
	fmt.Println(s)
	fmt.Println("---end---")
}

func (s Monster) GetNum(n1, n2 int) int {
	return n1 + n2
}

func (s Monster) Set(name string, age int, score float32, sex string) {
	s.Name = name
	s.Age = age
	s.Score = score
	s.Sex = sex
}
