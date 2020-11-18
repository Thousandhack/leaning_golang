package main

import "fmt"

// 自定义类型和类型别名
// type 后面跟着是类型
type myInt int
type youInt = int   // 这个叫类型别名
// 类型别名编译完成后就不存在了

func main() {
	var n myInt
	var m youInt
	n = 100
	m = 100
	fmt.Println(n)
	fmt.Printf("%T\n", n)

	fmt.Println(m)
	fmt.Printf("%T\n", m)

	var c rune
	c = '中'
	fmt.Println(c)

}
