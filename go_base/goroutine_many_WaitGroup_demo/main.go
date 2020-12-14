package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

// waitGroup

func f() {
	rand.Seed(time.Now().UnixNano())  // 保证每次执行不一样
	for i := 0; i < 5; i++ {
		ret1 := rand.Int()
		ret2 := rand.Intn(10) // 表示  0<= x < 10
		fmt.Println(ret1, ret2)
	}
}

func f1(i int){
	defer wg.Done()  // 运行结束减 1
	time.Sleep(time.Microsecond*time.Duration(rand.Intn(300)))
	fmt.Println(i)
}

var wg sync.WaitGroup

func main()  {
	//f()
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go f1(i)
	}
	// 知道goroutine结束了
	fmt.Println("main")
	wg.Wait()
}
