package main

import "fmt"

func test() {
	// 使用defer + recover 来捕获和处理异常
	defer func() {
		err := recover() //内置函数recover可以捕获到异常
		// if err::= recover(); err != nil {   // 另外一种写法
		if err != nil { // 说明捕获到异常或错误
			fmt.Println("err=", err)
			fmt.Println("错误处理方式")
		}
	}()
	num1 := 10
	num2 := 0
	res := num1 / num2
	fmt.Println("res=", res)
}

func main() {
	//
	test()
	// 下面的代码正常指向
	fmt.Println("test() 后运行的")
}
