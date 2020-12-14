package main

import (
	"fmt"
	"sync"
)

//var a []int
var b chan int // 需要指定通道中元素的类型
var wg sync.WaitGroup

func noBuffChannel() {
	fmt.Println(b)     // nil
	b = make(chan int) // 不带缓存区通道的初始化
	wg.Add(1)
	go func() {
		defer wg.Done()
		x := <-b
		fmt.Println("后台goroutine从通道b中渠道了...", x)
	}()
	b <- 10 // 如果只要发送会卡住了，报错
	fmt.Println(b)
	fmt.Println("10发送到通道b中了...")
	wg.Wait()
}

func buffChannel() {
	fmt.Println(b)         // nil
	b = make(chan int, 16) // 带缓存区通道的初始化
	b <- 10                // 如果只要发送会卡住了，报错
	fmt.Println(b)
	fmt.Println("10发送到通道b中了...")
	x := <-b
	fmt.Println("从通道b中渠道了...", x)

	b <- 20 // 如果只要发送会卡住了，报错
	fmt.Println(b)
	fmt.Println("20发送到通道b中了...")
	y := <-b
	fmt.Println("从通道b中渠道了...", y)
	close(b) // 关闭通道
}

func main() {
	//noBuffChannel()
	buffChannel()
}
