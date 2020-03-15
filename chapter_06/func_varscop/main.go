package main

import "fmt"

var age int = 50
var Name string = "zero"

// 函数
func test() {
	// age 和 Name 的作用域只在test函数内
	age := 10
	Name := "hsz"
	fmt.Println("age=", age)
	fmt.Println("Name=", Name)
}

func main() {
	test()
	fmt.Println("age=", age)   // 打印是test() 函数外的age变量
	fmt.Println("Name=", Name) // 打印是test()函数外的Name变量
	// 最后输出结果为：
	// age= 10
	// Name= hsz
	// age= 50
	// Name= zero
}
