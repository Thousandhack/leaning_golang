package main

import "fmt"

// 函数

// 函数的定义

// 没有返回值的
func sub(x int, y int) {
	fmt.Println(x - y)
}

// 有返回值的
// 函数的返回可以返回也可以不返回
func sum(x int, y int) int {
	return x + y
}

// 命名的返回值就相当于在函数中声明一个变量
// ret 不要再次定义
func sumTwo(x int, y int) (ret int) {
	ret = x + y
	return ret
}

// 多个参数返回值
func manyValue(x int, y int) (int ,int) {
	return x, y
}

func main() {
	sub(10, 5)

	ret := sum(5, 6)
	fmt.Println(ret)

	res := sumTwo(9, 4)
	fmt.Println(res)

	a, b := manyValue(5, 6)
	fmt.Println(a,b)
}
