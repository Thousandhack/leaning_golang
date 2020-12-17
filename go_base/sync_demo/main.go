package main

import (
	"fmt"
	"sync"
)

// 互斥锁的应用
var x = 0
var wg sync.WaitGroup
var lock sync.Mutex

func add(){
	for i:=0; i<50000;i++{
		lock.Lock()  // 加速后就不会造成竞争拿值
		x = x+1
		lock.Unlock()
	}
	wg.Done()
}



func main()  {
	wg.Add(2)
	go add()
	go add()
	wg.Wait()
	fmt.Println(x)
}
