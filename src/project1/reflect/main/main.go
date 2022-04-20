package main

import (
	"fmt"
	"reflect"
)

func main() {
	var num int = 100
	reflectTest01(num)

	var stu = Student{
		Name: "tome",
		Age:  20,
	}
	reflectTest02(stu)

	reflect01(&num)
	fmt.Println("num=", num)
}

//通过反射改变int型变量的值
//改变变量的值需要传入指针
func reflect01(b interface{}) {
	rVal := reflect.ValueOf(b)
	fmt.Printf("rVal kind=%v\n", rVal.Kind())
	//通过Elem()拿到对应的指针
	//等价于n int=*num
	rVal.Elem().SetInt(20)
}

func reflectTest01(b interface{}) {
	//1.通过反射获取reflect.TypeOf
	rType := reflect.TypeOf(b)
	fmt.Println("rType=", rType)
	//2.获取reflect.Value
	rVal := reflect.ValueOf(b)
	//获取reflect.Value的真实值
	//方式一：通过Int()函数
	n2 := 2 + rVal.Int()
	fmt.Println("n2=", n2)
	fmt.Printf("rVal=%v, rVal Type=%T\n", rVal, rVal)
	//方式二：
	iV := rVal.Interface()
	//接口通过断言转换类型
	num2 := iV.(int)
	fmt.Println("num2=", num2)
}

type Student struct {
	Name string
	Age  int
}

//演示对结构体的反射
func reflectTest02(b interface{}) {
	//1.通过反射获取reflect.TypeOf
	rType := reflect.TypeOf(b)
	fmt.Println("rType=", rType)
	//2.获取reflect.Value
	rVal := reflect.ValueOf(b)

	//3.获取变量对应的kind
	//(1) rVal.Kind() ==>
	kind1 := rVal.Kind()
	//(2) rType.Kind() ==>
	kind2 := rType.Kind()
	fmt.Printf("kind1=%v,kind2=%v\n", kind1, kind2)

	iV := rVal.Interface()
	fmt.Printf("iv=%v, iv Type=%T\n", iV, iV)
	stu, ok := iV.(Student)
	if ok {
		fmt.Printf("stu.Name=%v\n", stu.Name)
	}
}
