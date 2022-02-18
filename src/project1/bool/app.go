package main

import (
	"fmt"
	"unsafe"
)

func main() {
	var b = false
	fmt.Println("b=", b, ",占用空间:", unsafe.Sizeof(b))

	var age int = 40

	if age > 30 {
		fmt.Println("ok1")
	}

	if !(age > 30) {
		fmt.Println("ok2")
	}
}
