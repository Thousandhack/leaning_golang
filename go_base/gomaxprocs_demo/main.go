package main

import (
	"fmt"
	"runtime"
	"sync"
)

// GOMAXPROCS
var wg sync.WaitGroup


func a()  {
	defer wg.Done()
	for i :=0; i < 10; i ++ {
		fmt.Printf("A:%d\n",i)
	}
}


func b()  {
	defer wg.Done()
	for i :=0; i < 10; i ++ {
		fmt.Printf("B:%d\n",i)
	}
}

func main()  {
	runtime.GOMAXPROCS(2) // 默认cpu逻辑核心数，默认跑满整CPU
	wg.Add(2)
	go a()
	go b()
	// 结果会是一种混合打印A B
	wg.Wait()
}