package main

import "fmt"

func main() {
	// golang 字符类型的例子
	var c1 byte = 'a'
	var c2 byte = '0'
	// 直接输出byte的值的时候，直接输出相应的码值
	fmt.Println("c1:", c1)
	fmt.Println("c2:", c2)
	// 如果我们希望输出对于字符，需要格式化输出
	// 格式化输出使用的是Printf
	// 注意Printf打印不自动换行
	fmt.Printf("c1=%c ", c1)
	fmt.Printf("c1=%c\n", c1)

	// 第一个输出北  第二个输出北对应的码值
	var c3 int = '北'
	fmt.Printf("c3=%c c3=%d\n", c3, c3)

	// 使用相应的码值可以输出对于的字符
	var c4 int = 120
	fmt.Printf("c4=%c\n", c4)

	// 输出和的另外一个ascii码
	var n1 = 'a' + 10
	fmt.Println("n1=", n1) // 输出对应的大小
	fmt.Printf("n1=%c", n1)  // 输出ascii字符
}
