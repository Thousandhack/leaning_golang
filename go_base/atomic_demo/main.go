package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)

// 原子操作
var x int64
var wg sync.WaitGroup
//var lock sync.Mutex

func add(){
	//lock.Lock()
	//x ++
	//lock.Unlock()
	//wg.Done()

	// 与上面的锁的效果一致
	atomic.AddInt64(&x,1)
	wg.Done()
}

func main(){
	wg.Add(1000000)
	for i:=0;i<1000000;i++{
		go add()
	}
	wg.Wait()
	fmt.Println(x)
}
