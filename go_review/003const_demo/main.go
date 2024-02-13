package main

import "fmt"

const pi = 3.1415926

// 批量声明常量
const (
	statusOk = 200
	notFound = 404
)

// 只写一个数据常量，后面没有赋值默认和上一行一致
const (
	n1 = 100
	n2
	n3
)

// iota
// 在const关键字出现时将被重置为0。const中每新增一行常量声明将使iota计数一次(iota可理解为const语句块中的行索引)
// 使用iota能简化定义，在定义枚举时很有用。
const (
	a1 = iota // 0
	a2        // 1
	a3        // 2
	a4        // 3
	a5        // 4
)

func main() {
	fmt.Print(pi)
	fmt.Println(statusOk)
	fmt.Println(notFound)
	fmt.Println(n1)
	fmt.Println(n2)
	fmt.Println(n3)

	fmt.Println(a1)
	fmt.Println(a2)
	fmt.Println(a3)
	fmt.Println(a4)
	fmt.Println(a5)
}
