package main

import "fmt"

// 一个test函数
func test(n1 int){
	n1 += 1
	fmt.Println("n1=",n1)
}

// 一个get函数
func getSum(n1 int, n2 int) int {
	//
	sum := n1 + n2
	return sum
}

// 返回两个值的和和查
func getSumAndSub(a int, b int) (int, int){
	sum := a + b
	sub := a - b
	return sum,sub
}

func main(){
	// 调用 函数
	n1 := 10
	test(n1)
	get_sum := getSum(1, 2)
	fmt.Println("n1=",n1)
	fmt.Println("get_sum",get_sum)

	// 返回多个值
	re1,re2 := getSumAndSub(3,2)
	fmt.Println(re1)
	fmt.Println(re2)

	// 希望忽略某个返回值，用 _ 符号表示占位符
	_,re3 := getSumAndSub(5,4)
	fmt.Println(re3)
}