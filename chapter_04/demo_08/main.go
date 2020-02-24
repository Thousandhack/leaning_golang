package main

import "fmt"

func main() {
	// 位运算例子
	var a int = 1 >> 2 // 相当于  1 / 2  在go中得0
	var b int = -1 >> 2
	var c int = 1 << 2
	var d int = -1 << 2
	// a,b,c,d结果是什么
	fmt.Println("a=", a)
	fmt.Println("b=", b)
	fmt.Println("c=", c)
	fmt.Println("d=", d)


	fmt.Println(2&3)
	fmt.Println(2|3)
	fmt.Println(13&7)
	fmt.Println(5|4)
	fmt.Println(-3^3)

}
