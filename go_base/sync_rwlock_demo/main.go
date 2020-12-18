package main

import (
	"fmt"
	"sync"
	"time"
)

// 读写锁的应用
// 读用读锁
// 写用读写锁

var (
	x int64
	wg sync.WaitGroup
	//lock sync.Mutex
	rwLock sync.RWMutex
)


func write(){
	rwLock.Lock()   // 加写锁
	x = x + 1
	fmt.Println(x)
	time.Sleep(5*time.Microsecond)  // 程序耗时
	rwLock.Unlock()  // 解写锁
	wg.Done()

}


func read(){
	rwLock.RLock()   // 加写锁
	fmt.Println("read:",x)
	time.Sleep(10*time.Microsecond)  // 程序耗时
	rwLock.RUnlock()  // 解写锁
	wg.Done()
}

func main(){
	start := time.Now()
	for i :=0; i < 10;i++{
		wg.Add(1)
		go write()
	}

	for i :=0; i < 1000;i++{
		wg.Add(1)
		go read()
	}
	wg.Wait()
	end := time.Now()
	fmt.Println(end.Sub(start))

}