package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	rand.Seed(time.Now().Unix())
	n := rand.Intn(100) + 1
	fmt.Println(n)

	var count = 0
	for {
		n = rand.Intn(100) + 1
		fmt.Println(n)
		count++
		if n == 99 {
			break
		}
	}
	fmt.Println("生成99一共使用了", count)

label2:
	for i := 0; i < 4; i++ {
		//label1:
		for j := 0; j < 10; j++ {
			if j == 2 {
				//break
				//break label1
				break label2
			}
			fmt.Println("j=", j)
		}
	}
}
