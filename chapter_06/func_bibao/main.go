package main

import "fmt"

// 累加器
func Addupper() func(int) int {
	// 匿名函数与外面的n构成一个闭包
	var n int = 10
	var str = "hello"
	return func(x int) int {  // 返回是一个函数   匿名函数引用到外面变量n
		n = n + x
		str = str + string("a")
		fmt.Println("str=",str)
		return n
	}

}

func main() {
	//
	f := Addupper()
	fmt.Println(f(1))  // 11
	fmt.Println(f(2))  // 13

}
