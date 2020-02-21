package main

import "fmt"

func main() {
	// 演示 & 和 * 的使用
	a := 200
	fmt.Println("a的地址=",&a)

	var ptr *int = &a
	fmt.Println("ptr指向的值是=",*ptr)
}