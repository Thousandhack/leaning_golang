package main

import (
	"fmt"
	"sync"
)

// 启动单个goroutine
var wg sync.WaitGroup


func hello(i int) {
	defer wg.Done()  // 运行结束减 1
	fmt.Println("hello",i)
}

// 程序启动之后会创建一个主 goroutine 去执行
func main() {

	for i:=0;i<100;i++{
		wg.Add(1)  // 加1
		go hello(i)  // 开启一个单独的goroutine去执行hello函数（任务）
	}
	fmt.Println("main")
	wg.Wait()
}
