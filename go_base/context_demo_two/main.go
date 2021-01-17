package main

import (
    "fmt"
    "sync"
    "time"
)

// 为什么需要context?

var wg sync.WaitGroup
var exitChan = make(chan struct{})
//var exitChan = make(chan bool, 1) // make进行初始化,这种方法对应下面的注释的true

func f() {
    defer wg.Done()
FORLOOP:
    for {
        fmt.Println("TWO")
        time.Sleep(time.Millisecond * 500)
        select {
        case <-exitChan:
            break FORLOOP
        default:

        }
    }
}

func main() {
    wg.Add(1)
    go f()
    time.Sleep(time.Second * 2)
    // 通过修改全局变量的变化
    // 如果通知子goroutine退出
    exitChan <- struct{}{}
    //exitChan <- true

    wg.Wait()

}
