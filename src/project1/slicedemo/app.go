package main

import "fmt"

func main() {
	var intArr [5]int = [...]int{1, 2, 3, 4, 5}
	//声明定义一个切片
	//1.slice是切片的名称
	//2.intArr[1:3]表示引用数组intArr数组下标范围[1,3)的元素
	//省略startIndex默认为0
	//slice := intArr[:3]
	//省略endIndex默认为len(intArr)
	//slice:=intArr[1:]
	//省略startIndex和endIndex默认为整个数组
	//slice:=intArr[:]
	slice := intArr[1:3]
	slice[0] = 34
	fmt.Println("intArr=", intArr)
	fmt.Println("slice=", slice)
	//cap()获取切片的容量，切片的容量是可以动态变化的
	fmt.Println("slice=", slice, "len=", len(slice), "capacity=", cap(slice))

	fmt.Printf("intArr[1]的地址：%p\n", &intArr[1])
	fmt.Printf("splice[0]的地址：%p\n", &slice[0])
	fmt.Printf("splice的地址：%p\n", &slice)

	//定义切片的第二种方式
	//参数一：切片的类型
	//参数二：切片的长度
	//参数三：切片的容量，可选
	var slice2 []int = make([]int, 8, 10)
	slice2[1] = 20
	slice2[2] = 30
	fmt.Println(slice2)
	fmt.Println("slice2 len:", len(slice2))
	fmt.Println("slice2 cap:", cap(slice2))

	fmt.Println()
	//切片使用的第三种方式
	var slice3 []string = []string{"tom", "jack", "mary"}
	fmt.Println(slice3)
	fmt.Println("slice3 len:", len(slice3))
	fmt.Println("slice3 cap:", cap(slice3))
	//append函数追加元素
	slice4 := append(slice3, "hello", "world")
	for i, v := range slice3 {
		fmt.Printf("i=%v,v=%v\n", i, v)
	}
	fmt.Println()
	for i, v := range slice4 {
		fmt.Printf("i=%v,v=%v\n", i, v)
	}
	slice5 := append(slice3, slice4...)
	fmt.Println()
	for i, v := range slice5 {
		fmt.Printf("i=%v,v=%v\n", i, v)
	}

	//切片的拷贝
	slice6 := []int{1, 2, 3, 4, 5}
	slice7 := make([]int, 10)
	//将slice6拷贝到slice7
	//参数类型都是切片，数据都是独立的，相互不影响
	copy(slice7, slice6)
	fmt.Println()
}
