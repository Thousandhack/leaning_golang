package main

import "fmt"

// 案例2
type myFunType func(int, int) int // 这时myFun就是func(int,int) int 类型

func getSum(n1 int, n2 int) int {
	return n1 + n2
}

func myFun(funvar myFunType, num1 int, num2 int) int {
	return funvar(num1, num2)
}

func main() {

	type myInt int // 给int取别名 ，在go中myInt和int虽然都是int类型，但是go认为myInt和int还是不同的类型
	var num1 myInt
	var num2 int

	fmt.Printf("num1类型%T,num2类型%T\n", num1, num2) // 打印：num1类型main.myInt,num2类似int

	res := myFun(getSum, 500, 600)
	fmt.Println(res)  \\ 1100

}
