package main

import "fmt"

type myFunType func(int, int) int

func main() {
	var n1, n2 = 10, 20
	swap(&n1, &n2)
	fmt.Printf("n1=%v,n2=%v\n", n1, n2)

	fmt.Printf("%v\n", sumFloat(1, 2))

	//可变参数的使用
	sum := sum1(1, 2, 3, 4, 5)
	fmt.Printf("sum=%v\n", sum)

	//自定义变量类型
	//在go重myInt和int虽然都是int类型，但是go认为myInt和int是不同的类型
	//type myInt int
	//var num1 myInt = 40
	//var num2 int
	//num2 = int(num1)
	//fmt.Println("num1=", num1, "num2=", num2)
	//
	////函数变量
	//result := getSum(sum, 10, 20)
	//fmt.Printf("result=%v\n", result)
	//
	//a, b := getSumAndSub(20, 10)
	//fmt.Printf("a=%v,b=%v\n", a, b)

	//a := getSumAndSub
	//fmt.Printf("a的类型:%T, getSumAndSub的类型: %T\n", a, getSumAndSub)
	//res, _ := a(1, 2)
	//fmt.Printf("res=%v\n", res)

	//res1, res2 := getSumAndSub(1, 2)
	//fmt.Printf("res1=%v,res2=%v\n", res1, res2)
	//
	////只拿差
	//_, res3 := getSumAndSub(3, 2)
	//fmt.Printf("res3=%v\n", res3)
	//
	//var n int = 20
	//fmt.Printf("main->n的地址：%v\n", &n)
	//test01(&n)
	//fmt.Printf("n=%v\n", n)

}

//计算两个数的和和差
//返回值命名
func getSumAndSub(n1 int, n2 int) (sum int, sub int) {
	sum = n1 + n2
	sub = n1 - n2
	return
}

//通过指针改变参数的值
func test01(n *int) {
	fmt.Printf("test01->n的地址：%v\n", &n)
	*n = *n + *n
	fmt.Printf("test01 -> %v\n", *n)
}

//求两个数之和
func sum(n1 int, n2 int) int {
	return n1 + n2
}

//调用传入的函数
func getSum(funvar myFunType, num1 int, num2 int) int {
	return funvar(num1, num2)
}

//可变参数
func sum1(n1 int, args ...int) int {
	sum := n1
	for i := 0; i < len(args); i++ {
		sum += args[i]
	}
	return sum
}

func sumFloat(f1, f2 float32) float32 {
	fmt.Printf("f1 type:%T\n", f1)
	return f1 + f2
}

//交换两个数的值
func swap(n1 *int, n2 *int) {
	var t = *n1
	*n1 = *n2
	*n2 = t
}
