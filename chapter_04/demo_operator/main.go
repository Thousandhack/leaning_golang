package main

import "fmt"

func main() {
	// 运算符举例
	var a int = 21
	fmt.Println(+1)
	fmt.Println(-1)
	fmt.Println(5 + 5)
	fmt.Println(5 - 3)
	fmt.Println(3 * 4)
	fmt.Println(6 / 3)
	fmt.Println(6 % 4)

	a--
	fmt.Printf(" a 的值为 %d\n", a)
	a = 21 // 为了方便测试，a 这里重新赋值为 21
	a++
	fmt.Printf(" a 的值为 %d\n", a)
}
