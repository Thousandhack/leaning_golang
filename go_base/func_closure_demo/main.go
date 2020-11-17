package main

import (
	"fmt"
	"strings"
)

// 接收一个int类型的参数，返回一个int类型的结果
func adder() func(int) int {
	var x = 100
	return func(y int) int {
		x += y
		return x
	}
}

func f1(f func()) {
	fmt.Println("this is f1")
	f()
}

func f2(x, y int) {
	fmt.Println("this is f2")
	fmt.Println(x + y)
}

// 定义一个函数对f2进行包装
func f3(f func(int, int), m int, n int) func() {
	tmp := func() {
		f(m, n)
		fmt.Println("hh")
	}

	return tmp
}

// 给文件加后缀的函数
func makeSuffixFunc(suffix string) func(string) string {
	return func(name string) string {
		if !strings.HasSuffix(name, suffix) {
			return name + suffix
		}
		return name
	}
}

// 两个数的加减运算的值
func calc(base int) (func(int) int, func(int) int) {
	add := func(i int) int {
		base += i
		return base
	}

	sub := func(i int) int {
		base -= i
		return base
	}
	return add, sub
}

func main() {
	// 闭包指的是一个函数和与其相关的引用环境组合而成的实体。简单来说，闭包=函数+引用环境。
	// 这个函数包含外部作用域的变量
	// 闭包的底层原理：
	// 1. 函数可以作为返回值
	// 2. 函数内部查找变量的顺序，先在自己内部找，找不到往外层找
	// 闭包的例子
	ret := adder()
	ret2 := ret(200)
	fmt.Println(ret2)

	// 应用例子
	// f2 传到 f1 里面去执行
	res1 := f3(f2, 100, 200)
	fmt.Printf("%T\n", res1)
	f1(res1)

	// 给文件名加.jpg
	jpgFunc := makeSuffixFunc(".jpg")
	// 给文件名加.txt
	txtFunc := makeSuffixFunc(".txt")
	fmt.Println(jpgFunc("test")) //test.jpg
	fmt.Println(txtFunc("test")) //test.txt

	//
	c1, c2 := calc(10)
	fmt.Println(c1(1), c2(2)) //11 9
	fmt.Println(c1(3), c2(4)) //12 8
	fmt.Println(c1(5), c2(6)) //13 7

}
