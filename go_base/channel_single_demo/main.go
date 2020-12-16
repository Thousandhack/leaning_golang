package main

import (
	"fmt"
	"sync"
)

// 单项通道

var wg sync.WaitGroup
var once sync.Once

func f1(ch1 chan<- int) {
	defer wg.Done()
	for i := 0; i < 100; i++ {
		ch1 <- i
	}
	close(ch1) // 通道关闭之后不能再继续写入，但是可以读出
}

// ch1 <-chan int, ch2 chan<- int
// 以上表示只能从ch1 取值，ch2 只能放值
func f2(ch1 <-chan int, ch2 chan<- int) {
	defer wg.Done()
	for {
		x, ok := <-ch1
		if !ok {
			break
		}
		ch2 <- x * x
	}
	// 确保只执行一次
	once.Do(func() {
		close(ch2)
	})
}

func main() {
	a := make(chan int, 100)
	b := make(chan int, 100)
	wg.Add(3)
	go f1(a)
	// 分给两个goroutine去消费
	go f2(a, b)
	go f2(a, b)
	wg.Wait()
	for ret := range b {
		fmt.Println(ret)
	}
}
