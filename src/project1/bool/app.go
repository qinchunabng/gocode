package main

import (
	"fmt"
	"unsafe"
)

func main() {
	var b = false
	fmt.Println("b=", b, ",占用空间:", unsafe.Sizeof(b))
}
