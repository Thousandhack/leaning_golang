package main

import (
	"fmt"
	"time"
)

// goroutine

func hello(i int) {
	fmt.Println("hello",i)
}

// 程序启动之后会创建一个主 goroutine 去执行
func main() {
	//  下面谁先跑是不一定的
	for i:=0;i<100;i++{
		//go hello(i)  // 开启一个单独的goroutine去执行hello函数（任务）
		// 启动goroutine需要花费时间的
		go func(i int) {
			fmt.Println(i,"匿名func")  // 用的是函数参数的i ,不是外面的i
		}(i)
	}
	fmt.Println("main")
	time.Sleep(time.Second)
	// main 函数结束了 由main函数启动的goroutine也都结束了
}
