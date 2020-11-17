package main

import "fmt"



func main() {
	// 匿名函数
	// 没有名字的函数
	// 匿名函数一般用于函数内部，因为函数内部没有办法声明带有名字的函数
	var f1 = func (x int, y int) int {
		return x + y
	}
	res := f1(5,6)
	fmt.Println(res)

	// 如果只是调用一次的函数，还可以简写成如下：
	func(x,y int){
		fmt.Printf("hhh")
		//res := x+y
		fmt.Println(x+y)
	}(100,200)

}
