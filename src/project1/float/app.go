package main

import "fmt"

func main() {
	//演示小数类型的使用
	var price float32 = 89.12
	fmt.Println("price=", price)

	var num1 float32 = -0.00089
	var num2 float64 = -7809666.09
	fmt.Println("num1=", num1, "num2", num2)

	var num3 = 0.001
	num3, num4 := 0.451, 0.7845
	//默认是是float64
	fmt.Printf("num2 type: %T\n", num3)
	fmt.Println("num3=", num3, "num4=", num4)
	fmt.Printf("num4 type: %T", num4)
	//科学计数法
	num5 := 5.1234e2
	num6 := 5.1234e-2
	fmt.Println("num5=", num5, "num6=", num6)
}
