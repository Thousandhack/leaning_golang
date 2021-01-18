package main

import (
    "context"
    "fmt"
    "sync"
    "time"
)

// 为什么需要context?
// 下面这个例子是使用context来解决通知子goroutine关闭的问题
// 可以用于请求的超时
var wg sync.WaitGroup

func f(ctx context.Context) {
    defer wg.Done()
FORLOOP:
    for {
        fmt.Println("TWO")
        time.Sleep(time.Millisecond * 500)
        select {
        case <-ctx.Done():
            break FORLOOP
        default:

        }
    }
}

func main() {
    ctx, cancel := context.WithCancel(context.Background())
    wg.Add(1)
    go f(ctx)
    time.Sleep(time.Second * 2)
    // 通过修改全局变量的变化
    // 如果通知子goroutine退出
    cancel()

    wg.Wait()

}
