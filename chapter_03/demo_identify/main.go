package main

import "fmt"

// golang中标识符的使用
func main() {
	// Golang 中严格区分大小写
	var num int = 10
	var Num int = 20
	fmt.Printf("num=%v\nNum=%v\n", num, Num)

	// 标识符不能包含空格
	// var ab c int = 30

	// _ 是空标识符，用于占位
	// var _ int = 40 // 这样是错误的

	// 语法可以通过，但是要求不能使用
	var int int = 90
	fmt.Println(int)
}
