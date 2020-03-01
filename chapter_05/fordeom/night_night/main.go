package main

import "fmt"

func main() {
	// 打印九九乘法表的格式
	for i := 1; i <= 9; i++{
		for j := 1; j <= i; j++{
			fmt.Printf("%v*%v=%v\t",i,j,i*j)
		}
		fmt.Printf("\n")
	}
}