package main

import "fmt"

func main() {
	ch := make(chan int, 1)
	// 因为缓存只设置了1 ，如果存入一个值
	// 那么第二次循环就到另外一个case
	for i := 0; i < 10; i++ {
		select {
		case x := <-ch:
			fmt.Println(x)
		case ch <- i:
		}

	}
}