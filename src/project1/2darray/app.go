package main

import "fmt"

func main() {
	//二维数组
	//4行6列
	//[0,0,0,0,0,0]
	//[0,0,0,0,0,0]
	//[0,0,0,0,0,0]
	//[0,0,0,0,0,0]
	var arr [4][6]int
	arr[1][2] = 1
	fmt.Println(arr)

	for i := 0; i < len(arr); i++ {
		fmt.Println(arr[i])
	}

	for i := 0; i < len(arr); i++ {
		for j := 0; j < len(arr[i]); j++ {
			fmt.Print(arr[i][j], " ")
		}
		fmt.Println()
	}

	var arr2 [2][3]int = [2][3]int{{1, 2, 3}, {4, 5, 6}}
	fmt.Println(arr2)
}
