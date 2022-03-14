package main

import "fmt"

func main() {
	var c = Circle{20}
	res := c.Area()
	fmt.Println("res=", res)
}

type Circle struct {
	Radius float64
}

func (c *Circle) Area() float64 {
	return 3.14 * c.Radius * c.Radius
}
