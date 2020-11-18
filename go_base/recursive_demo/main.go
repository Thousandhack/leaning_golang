package main

import "fmt"

// 阶乘
func funcDemo(n uint64) (result uint64) {
	if n <= 1 {
		return 1
	}
	result = n * funcDemo(n-1)
	return result
}

// 走台阶，n个台阶，要么走一步，要么走两步，有多少种走法
// 解题思路：要么剩一个台阶，要么剩两个台阶
func funcTaiJie(n int) int {
	if n <= 2 {
		return n
	}
	return funcTaiJie(n-2) + funcTaiJie(n-1)
}

func main() {
	ret := funcDemo(5)
	fmt.Println(ret)

	res := funcTaiJie(4)
	fmt.Println(res)
}
