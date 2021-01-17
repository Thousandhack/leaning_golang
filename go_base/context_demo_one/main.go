package main

import (
    "fmt"
    "sync"
    "time"
)

// 为什么需要context?

var wg sync.WaitGroup
var notify bool

func f(){
    defer wg.Done()
    for {
        fmt.Println("hsz")
        time.Sleep(time.Millisecond * 500)
        if notify {
            break
        }
    }
}

func main(){
    wg.Add(1)
    go f()
    time.Sleep(time.Second * 2)
    // 通过修改全局变量的变化
    // 如果通知子goroutine退出
    notify = true
    wg.Wait()

}
