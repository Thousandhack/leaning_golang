package main

import "fmt"

func main() {
	// 浮点类型的使用
	var price float32 = 19.23
	fmt.Println("price:", price)

	// 尾数部分可能丢失，造成精度损失
	var num3 float32 = -123.2345623456
	var num4 float64 = -123.2345623456
	fmt.Println("num3:", num3)
	fmt.Println("num4:", num4)
	// 打印结果如下：
	// num3: -123.234566
	// num4: -123.2345623456

	// 测试浮点型的默认精度
	var b = 1.11
	fmt.Printf("b的类型为%T\n", b)

	// 科学计算法
	num5 := 5.12e2
	fmt.Println("num5:", num5)

}
