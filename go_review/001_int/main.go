package main

import "fmt"

func main() {
	// 十进制
	var a int = 10
	fmt.Printf("%d \n", a) // 10
	fmt.Printf("%b \n", a) // 1010  占位符%b表示二进制

	// 八进制  以0开头
	var b int = 077
	fmt.Printf("%o \n", b) // 77

	// 十六进制  以0x开头
	var c int = 0xff
	fmt.Printf("%x \n", c) // ff
	fmt.Printf("%X \n", c) // FF

	// 进制转换
	var d = 101
	fmt.Printf("%d\n", d)
	fmt.Printf("%b\n", d)
	fmt.Printf("%o\n", d)
	fmt.Printf("%x", d)
}
