package main

import "fmt"

func main() {
	//管道可以声明为只读或只写
	//默认情况管道是双向的，可读可写

	//声明为只写
	var chan1 chan<- int
	chan1 = make(chan int, 3)
	chan1 <- 20
	//num := <-chan1

	fmt.Println("chan1=", chan1)

	//声明为只读
	//var chan2 <-chan int
	////chan2 = make(chan int, 2)
	//num3 := <-chan2
	//fmt.Println("num3=", num3)

	//应用场景实例
	var ch chan int
	ch = make(chan int, 10)
	exitChan := make(chan struct{}, 2)
	go send(ch, exitChan)
	go recv(ch, exitChan)

	var total = 0
	for _ = range exitChan {
		total++
		if total == 2 {
			break
		}
	}
	fmt.Println("结束....")
}

func send(ch chan<- int, exitChan chan struct{}) {
	for i := 0; i < 10; i++ {
		ch <- i
	}
	close(ch)
	var a struct{}
	exitChan <- a
}
func recv(ch <-chan int, exitChan chan struct{}) {
	for {
		v, ok := <-ch
		if !ok {
			break
		}
		fmt.Println(v)
	}
	var a struct{}
	exitChan <- a
}
