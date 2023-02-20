package main

import "fmt"

// fmt 占位符

func main() {
	var n = 100
	// 查看类型
	fmt.Printf("%T\n", n) // 类型
	fmt.Printf("%v\n", n) // 值
	fmt.Printf("%b\n", n) // 二进制
	fmt.Printf("%d\n", n) // 十进制
	fmt.Printf("%o\n", n) // 八进制
	fmt.Printf("%x\n", n) //十六进制

	n2 := "100"
	fmt.Printf("%s\n", n2) // 打印字符串

}
