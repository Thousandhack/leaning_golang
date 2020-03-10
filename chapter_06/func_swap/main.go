package main

import "fmt"

func swap(n1 *int, n2 *int) { // 指针接收地址进行值的相应交换
	// 定义一个临时变量
	t := *n1
	*n1 = *n2
	*n2 = t

}

func main() {
	a := 10
	b := 20
	swap(&a, &b) // 传入的是地址
	fmt.Printf("a=%v,b=%v", a, b)

}
