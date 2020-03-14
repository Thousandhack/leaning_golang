package main

import "fmt"

func sum(n1 int, n2 int) int {
	// defer 的作用会让其本来要执行的语句暂时不执行，将要执行的语句压入到一个defer栈（独立栈）
	// 暂时不执行
	// 当函数执行完毕后，再执行defer栈中语句：先入后出执行
	defer fmt.Println("n1=", n1) // defer
	defer fmt.Println("n2=", n2) // defer
	res := n1 + n2
	fmt.Println("in_res=", res)
	return res
}

func main() {
	res := sum(5, 7)
	fmt.Println("out_res=", res)
	// 结果：
	// in_res= 12
	// n2= 7
	// n1= 5
	// out_res= 12
}
