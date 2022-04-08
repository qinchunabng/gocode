package main

import (
	"fmt"
	"runtime"
)

func main() {
	cpuNum := runtime.NumCPU()
	fmt.Println("cpuNum=", cpuNum)

	//设置cpu
	runtime.GOMAXPROCS(cpuNum - 1)
	fmt.Println("ok")
}
