package main

import "fmt"

func main() {
	var hens [6]float64
	fmt.Println(hens)
	fmt.Printf("hens的地址:%p,hens[0]的地址:%p,hens[1]的地址:%p,hens[2]的地址:%p\n", &hens, &hens[0], &hens[1], &hens[2])
	hens[0] = 3.0
	hens[1] = 4.0
	hens[2] = 5.0
	hens[3] = 6.0
	hens[4] = 7.0
	hens[5] = 8.0
	totalWeight := 0.0
	for i := 0; i < len(hens); i++ {
		totalWeight += hens[i]
	}
	averageWeight := fmt.Sprintf("%.2f", totalWeight/float64(len(hens)))
	fmt.Printf("totalWeight=%v,averageWeight=%v\n", totalWeight, averageWeight)

	//数组初始化
	var numsArray01 [3]int = [3]int{1, 2, 3}
	fmt.Println("numsArray01=", numsArray01)

	var numsArray02 = [3]int{1, 2, 3}
	fmt.Println("numsArray02=", numsArray02)

	var numsArray03 = [...]int{8, 9, 10}
	fmt.Println("numsArray03=", numsArray03)

	var numArr04 = [...]int{1: 800, 0: 900, 2: 299}
	fmt.Println("numArr04=", numArr04)

	stringArray05 := [...]string{1: "tome", 0: "jack", 2: "mary"}
	fmt.Println("stringArray05=", stringArray05)

	//数组的遍历方式
	for i, s := range stringArray05 {
		fmt.Printf("i=%v,s=%v\n", i, s)
	}

	//在go中数组的长度是数据类型的一部分，[3]int和[4]int是不同的类型
	arr := [3]int{1, 2, 3}
	testArray(arr)
	fmt.Println("arr=", arr)
	testArray02(&arr)
	fmt.Println("arr=", arr)
}

//go中数组是值类型，这里传入函数的是变量的副本
func testArray(arr [3]int) {
	arr[0] = 250
}

//用指针修改数组的值
func testArray02(arr *[3]int) {
	(*arr)[0] = 88
}
