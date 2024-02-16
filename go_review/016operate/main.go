package main

import "fmt"

// 运算符
func main() {

	age := 31
	// 逻辑运算符
	if age > 18 && age < 60 {
		fmt.Println("苦逼上班")
	} else if age < 18 || age >= 60 {
		fmt.Println("上线或退休")
	} else {
		fmt.Println("不存在")
	}

	// 取反
	isStu := false
	fmt.Println(isStu)
	fmt.Println(!isStu)

	// 位运算符  针对的是二进制
	// 按位与 & 两个都1才为1
	// 101
	// 010
	// 000   => 0
	fmt.Println(5 & 2)
	// 按位或 | 两个有1就为1
	// 101
	// 010
	// 111   => 7

	fmt.Println(5 | 2)
	// 按位或  两位不一样则为1
	// 101
	// 010
	// 111  =>  7
	fmt.Println(5 ^ 2)
	// 将二进制数左移指定位数
	// 101 => 1010 => 10
	fmt.Println(5 << 1)
	// 101 => 10 => 2
	fmt.Println(5 >> 1)

	// 赋值运算符

	var x int
	x = 10
	x += 1
	fmt.Println(x)
}
