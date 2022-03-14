package main

import (
	"fmt"
	"project1/factory/model"
)

func main() {
	var stu = model.NewStudent("Jack", 80)
	fmt.Println(*stu)
	fmt.Println("stu score:", stu.GetScore())
}
