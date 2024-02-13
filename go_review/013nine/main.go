package main

import "fmt"

func main() {
	// 9 9 乘法表
	for i := 1; i < 10; i++ {
		for j := 1; j <= i; j++ {
			fmt.Printf("%d*%d=%d\t", i, j, i*j)
		}
		fmt.Println()
	}

	for i := 1; i < 10; i++ {
		for j := i; j >= 1; j-- {
			fmt.Printf("%d*%d=%d\t", i, j, i*j)
		}
		fmt.Println()
	}
}
