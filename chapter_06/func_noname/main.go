package main

import "fmt"


var (
	// func1 为全局匿名函数
	Func1 = func(n1 int,n2 int) int {
		return n1 * n2
	}
)

func main() {
	// 在定义匿名函数时就直接调用
	// 使用匿名函数求两个数和
	res := func (n1 int,n2 int) int {
		return n1 + n2
	}(10,20)

	fmt.Println(res)

	// 将匿名函数func (n1 int n2 int) int 赋给 a 变量
	// 则 a 的数据类型就是函数类型 此时 我们可以通过a完成调用
	a := func (n1 int,n2 int) int {
		return n1 - n2
	}

	// 可以多次调用  多态？？？
	res2 := a(30,10)
	fmt.Println(res2)
	res3 := a(40,10)
	fmt.Println(res3)

	// 全局匿名函数的使用
	res4 := Func1(4,5)
	fmt.Println(res4)

}
