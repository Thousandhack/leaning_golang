package main

import "fmt"

// 有参数有返回值的函数
func f1(a int, b int) int {
	return a + b
}

// 没有参数没有返回值的函数
func f2() {
	fmt.Println("f2")
}

// 没有参数但有返回值的函数
func f3() int {
	return 3
}

// 返回值可以也可以不命名
func f4(x, y int) (ret int) {
	ret = x + y
	return
}

// 返回多个值
func f5() (int, int) {
	return 1, 2
}

// 参数类型简写
func f6(x, y int, m, n string, i, j bool) int {
	return x + y
}

// 可变长参数
func f7(x string, y ...int) {
	fmt.Println(x)
	fmt.Printf("%T\n", y)
	fmt.Println(y)
}

func main() {
	fmt.Println(f1(3, 4))
	f2()
	fmt.Println(f3())
	fmt.Println(f4(1, 2))
	fmt.Println(f5())
	f6(2, 3, "a", "b", true, false)
	f7("下雨了", 1, 2, 3, 4, 5, 6, 7, 8, 9)
	f7("下雨了")
}
