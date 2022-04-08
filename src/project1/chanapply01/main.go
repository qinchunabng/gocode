package main

import "fmt"

func main() {
	intChan := make(chan int, 1000)
	primeChan := make(chan int, 2000)
	//标识退出的管道
	exitChan := make(chan bool, 4)

	//开启一个协程，向intChan中放入1-8000个数
	go putNum(intChan)

	//开启四个协程，从intChan中取出数据，判断是否为素数，如果是，放入到primeChan
	for i := 0; i < 4; i++ {
		go primeNum(intChan, primeChan, exitChan)
	}

	//从exitChan取出四个数据，表明四个协程都已经完成
	go func() {
		for i := 0; i < 4; i++ {
			<-exitChan
		}
		//去除4个结果，关闭primeChan管道
		close(primeChan)
	}()

	//遍历输出结果
	for {
		res, ok := <-primeChan
		if !ok {
			break
		}
		//将结果输出
		fmt.Printf("素数=%d\n", res)
	}
	fmt.Println("main线程退出")
}

//向intChan中放入1-8000个数
func putNum(intChan chan int) {
	for i := 1; i <= 8000; i++ {
		intChan <- i
	}
	close(intChan)
}

//从intChan中取出数据，判断是否是素数，如果是就放入到primeChan
func primeNum(intChan chan int, primeChan chan int, exitChan chan bool) {
	var flag bool
	for {
		num, ok := <-intChan
		if !ok { //取不到数
			break
		}
		flag = true
		//判断num是不是素数
		for i := 2; i < num; i++ {
			if num%i == 0 {
				flag = false
				break
			}
		}
		if flag {
			//将这个素数放入到primeChan
			primeChan <- num
		}
	}
	fmt.Println("因为有个协程因为取不到数据，退出")
	exitChan <- true
}
