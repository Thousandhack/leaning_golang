package main

import "fmt"

// defer多用于函数结束之前释放资源（文件句柄、数据连接、socket连接）
func deferDemo(){
	fmt.Println("哈哈哈")
	defer fmt.Println("呵呵呵")  // defer 把它后面的语句延迟到函数即将返回的时候再执行
	defer fmt.Println("嘿嘿嘿")  // 一个函数中可以有多个defer语句，多个defer按照先进后出的顺序执行
	fmt.Println("嘻嘻嘻")
}


// Go语言中函数的return不是原子操作，在底层分为两步来执行
// 第一步： 返回值赋值
// 第二步： 真正的RET返回
// 函数中如果存在defer,那么defer执行的时间是在第一步和第二步之间
func deferDemoTwo() int {
	x := 5
	defer func(){
		x ++   // 修改的是x 不是返回值
	}()
	return x
}

func deferDemoThree()(x int){
	defer func(){
		x ++
	}()
	return 5  // 返回值=x
}

func calc(index string, a, b int) int {
	ret := a + b
	fmt.Println(index, a, b, ret)
	return ret
}

func main(){
	deferDemo()
	// 5
	res := deferDemoTwo()
	fmt.Println(res)

	// 6
	fmt.Println(deferDemoThree())

	// defer语句延迟调用的特性，所以defer语句能非常方便的处理资源释放问题。比如：资源清理、文件关闭、解锁及记录时间等。

	x := 1
	y := 2
	defer calc("AA", x, calc("A", x, y))
	x = 10
	defer calc("BB", x, calc("B", x, y))
	y = 20
}
