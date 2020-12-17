package main

import (
	"fmt"
	"math/rand"
	"time"
)

// channel
// 以上的情况不会造成死锁，只会阻塞

func sendNum(ch chan<- int){
	for {
		num := rand.Intn(100)
		ch <- num
		time.Sleep(5 * time.Second)
	}
}

func main() {
	ch := make(chan int, 1)
	go sendNum(ch)
	for {
		x,ok := <-ch  // 阻塞等4秒
		fmt.Println(x,ok)
		time.Sleep(time.Second)
	}

}
