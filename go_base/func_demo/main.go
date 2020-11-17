package main

import "fmt"

var x = 100 // 定义一个全局变量

// 函数是组织好的、可重复使用的、用于执行指定任务的代码块。
// 函数的可变参数与返回切片
// return也可以省略返回值
func funcDemo(title string, y ...int) (z []int) {
	fmt.Println(title)
	return y
}

func globalVarDemo() {
	// 函数的查找变量的顺序
	// 1. 先在函数内部查找
	// 2. 找不到就往函数的外面查找，一直找到全局

	fmt.Println(x)
}

func funcTypeDemo() {
	fmt.Println("Hello")
}

func funcTypeDemoTwo() int {
	return 10
}

func funcTypeDemoFour(a int, b int) int {
	return a + b
}

// 传进去的参数为函数类型 下面为有一个返回值的函数
func funcTypeDemoThree(x func() int) func(int,int) int  {
	z := x()
	fmt.Println("z = ",z)
	return funcTypeDemoFour
}

func main() {
	sliceDemo := make([]int, 3, 3)
	sliceDemo[0] = 3
	sliceDemo[1] = 6
	sliceDemo[2] = 9
	sliceDemo = append(sliceDemo, 100)
	fmt.Println(sliceDemo)
	res := funcDemo("Hello", sliceDemo[0], sliceDemo[1])
	fmt.Println(res)

	// 全局变量的例子
	// 函数内的变量，无法在函数外使用
	// 全局作用域
	// 函数作用域
		// LEGB原则
	// 代码块作用域 for
	globalVarDemo()

	// 函数也是一个变量类型
	// 高阶函数
	// 函数也是一种数据类型，既可以做参数也可以做返回值
	fmt.Printf("%T\n", funcTypeDemo)
	fTDT := funcTypeDemoTwo
	fmt.Printf("%T\n", fTDT)  // 数据类型： func() int

	ret := funcTypeDemoThree(funcTypeDemoTwo)
	fmt.Printf("%T\n",ret)   // 返回值为函数
}
