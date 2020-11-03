package main

import "fmt"

func main() {
	// 十进制转换为 二进制 八进制和十六进制
	var i1 = 101
	fmt.Printf("%d\n", i1)
	fmt.Printf("%b\n", i1)
	fmt.Printf("%o\n", i1)
	fmt.Printf("%x\n", i1)

	// 八进制 打印十进制
	i2 := 077
	fmt.Printf("%d\n", i2)

	// 十六进制 打印十进制
	i3 := 0x1234567
	fmt.Printf("%d\n", i3)

	// 二进制 打印十进制
	i4 := 0b11
	fmt.Printf("%d\n", i4)
	fmt.Printf("%T\n", i4)  // 变量的类型
}
