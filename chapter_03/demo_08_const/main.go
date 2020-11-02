package main

import "fmt"

// 单个声明常量
const pi = 3.1415926

// 批量声明常量
const (
	statusOK = 200
	notFund  = 404
)

const (
	n1 = 100
	n2
	n3
)

// iota    const 中每新增一行常量声明将使iota计数一次(+1)
// 类似枚举
const (
	a1 = iota  // 0
	a2 = iota  // 1
	a3 = iota  // 2
)

const (
	b1 = iota  // 0
	b2 = iota // 1
	_ = iota // 2
	b3 = iota // 3
)

const (
	c1 = iota // 0
	c2 = 100  // 100
	c3 = iota //  2
	)

const (
	d1,d2 = iota + 1, iota + 2   // 1     2
	d3,d4 = iota + 1 , iota + 2  //  2    3
)

func main() {
	fmt.Println(pi)
	fmt.Println(statusOK)
	fmt.Println(notFund)
	fmt.Println(n1, n2, n3)
	fmt.Println(a1, a2, a3)
	fmt.Println(b1, b2, b3)
	fmt.Println(c1, c2, c3)
	fmt.Println(d1, d2, d3,d4)
}
