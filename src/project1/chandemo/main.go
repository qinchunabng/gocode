package main

import "fmt"

func main() {
	iterateChan()

	//创建一个存放3个int的管道
	var intChan chan int
	intChan = make(chan int, 3)

	fmt.Printf("intChan 的值=%v, intChan本身的地址=%p\n", intChan, &intChan)

	//向管道写入数据
	intChan <- 10
	num := 110
	intChan <- num

	//关闭管道
	//管道关闭后不能像管道写数据
	//只能读数据
	close(intChan)
	num = <-intChan
	intChan <- 100

	//注意：当向管道写入数据时，不能超过管道的长度

	//查看管道的长度和容量
	fmt.Printf("channel len=%v cap=%v \n", len(intChan), cap(intChan))

	intChan <- 1
	//超过管道的长度会报错
	//intChan <- 2

	//从管道中读取数据
	//注意：如果在管道中没有数据的情况，再取数据就会报deadlock的错误
	var num2 = <-intChan
	fmt.Println("num2=", num2)
	fmt.Printf("channel len=%v cap=%v \n", len(intChan), cap(intChan))

	//可以不赋值给任何变量，直接丢弃掉
	<-intChan
	//num2 = <-intChan
	//num3 := <-intChan
	//num4 := <-intChan
	//fmt.Println(num2, num3, num4)

	testChan()
}

type Cat struct {
	Name string
	Age  int
}

func testChan() {
	allChan := make(chan interface{}, 3)

	allChan <- 10
	allChan <- "tom"
	cat := Cat{
		Name: "小花猫",
		Age:  1,
	}
	allChan <- cat

	<-allChan
	<-allChan

	newCat := <-allChan
	fmt.Printf("newCat=%T, newCat=%v\n", newCat, newCat)
	//下面这行代码是错误的，必须使用类型断言
	//fmt.Printf("newCat.Name=%v", newCat.Name)
	//需要是由类型断言
	a, ok := newCat.(Cat)
	if ok {

	}

	fmt.Printf("newCat.Name=%v", a.Name)
}

//遍历管道
func iterateChan() {
	intChan := make(chan int, 100)
	for i := 0; i < 100; i++ {
		intChan <- i * 2
	}
	close(intChan)

	//管道遍历只能用foreach
	//在遍历时必须先关闭管道，不然会出现deadlock异常
	for v := range intChan {
		fmt.Println("v=", v)
	}
}
