package main

import (
	"fmt"
	"unicode"
)

func main() {
	//编写代码分别定义一个整型、浮点型、布尔型、字符串型变量，使用fmt.Printf()搭配%T分别打印出上述变量的值和类型。
	var a int32
	a = 10
	var b float32
	b = 10
	var c bool
	c = true
	var d string
	d = "10"
	fmt.Printf("%T\n", a)
	fmt.Printf("%T\n", b)
	fmt.Printf("%T\n", c)
	fmt.Printf("%T\n", d)
	//编写代码统计出字符串"hello深圳南山小王子"中汉字的数量。
	// 1.依次拿到字符串中的字符
	// 2.判断当前这个字符串中的字符
	// 3.把汉字的次数累加得到最终结果
	s := "hello深圳南山小王子"
	count := 0
	for _, v := range s {
		if unicode.Is(unicode.Han, v) {
			count++
		}

	}
	fmt.Printf("有%v个汉字", count)

}
