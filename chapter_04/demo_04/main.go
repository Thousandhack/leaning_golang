package main

import "fmt"


func test() int {
	return 1
}

func main() {
	// // 赋值运算符的使用演示
	// var i int
	// i = 10 // 基本赋值
	// fmt.Println(i)
	// // 有两个变量，a和b ,要求将其进行交换，最终打印结果
	// // a = 9， b = 2 ===>  a = 2 b = 9
	// a := 9
	// b := 2
	// fmt.Printf("交换前的情况是a=%v,b=%v\n", a, b)
	// // 定义一个临时变量
	// t := a
	// a = b
	// b = t
	// fmt.Printf("交换后的情况是a=%v,b=%v\n", a, b)

	// // 复合赋值操作
	// a += 17
	// fmt.Println(a)


	// // 
	// var c int = 5
	// var d int 
	// d = c + 3  // 赋值运算的执行顺序是从右向左
	// fmt.Println(d)  // 8
	// // 赋值运算符的左边，只能是变量，右边可以是变量、表达式、常量
	// var e int
	// e = c
	// fmt.Println(e)   // 5
	// e = 8 + c * 2
	// fmt.Println(e)  // 18
	// e = test() + 90
	// fmt.Println(e)  // 91
	// e = 900
	// fmt.Println(e)  // 900

	// 有两个变量，a和b ,要求将其进行交换，但不允许使用中间变量，最终打印结果

	var a int = 6
	var b int = 5
	fmt.Printf("开始值：a=%v,b=%v\n",a,b)
	
	a = a + b 
	b = a - b
	a = a - b
	fmt.Printf("交换后：a=%v,b=%v\n",a,b)

}
