package main

import (
	"fmt"
	"project1/packagedemo/utils"
)

var Test = 10

func main() {
	result := utils.Calc(1.2, 1.3, '+')
	fmt.Printf("result=%.2f\n", result)
}
