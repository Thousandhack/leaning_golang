package main

import (
	"fmt"
)

func main() {
	// 字符串修改
	s2 := "白萝卜"
	s3 := []rune(s2) // 把字符串强制转成一个rune切片
	s3[0] = '红'
	fmt.Println(string(s3))

	// rune 和 byte 有区别
	c3 := "H"
	c4 := 'h'
	fmt.Printf("c3:%T c4:%T\n", c3, c4)
	fmt.Printf("%d\n", c4)

	// 类型转换
	n1 := 10 // int
	var f float64
	f = float64(n1)
	fmt.Println(f)
	fmt.Printf("%T\n", f)

}
