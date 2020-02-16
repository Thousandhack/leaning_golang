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
	// 除法使用的特点:
	// 说明： 如果运算的数都是整数，那么除后，去掉小数部，保留整数部分
	fmt.Println(10 / 4)

	var n1 float32 = 10 / 4
	fmt.Println(n1)

	// 如果我们希望保留小数部分，则需要有浮点数参与运算
	var n2 float32 = 10.0 / 4
	fmt.Println(n2)

	// 取模使用的特点：
	// 取模公式如下：a % b = a - a / b * b
	fmt.Println("10模3等于  ：", 10%3)     //  1
	fmt.Println("-10模3等于 ：", -10%3)   //  = -10 - (-10) / 3 * 3 = -10 - (-9) = -1
	fmt.Println("10模-3等于 ：", 10%-3)   //  = 10 - 10/(-3)*(-3) =   10 - 9 = 1
	fmt.Println("-10模-3等于：", -10%-3) //  = -10 - （-10）/(-3)*(-3) = -10 + 9 = -1

	a--
	fmt.Printf(" a 的值为 %d\n", a)
	a = 21 // 为了方便测试，a 这里重新赋值为 21
	a++
	fmt.Printf(" a 的值为 %d\n", a)
}
