package main

import "fmt"
import "math/rand"
import "time"

func main() {
	// 随机生成1-100的一个数，直到生成了99这个数，看看一共需要用多少次
	var conut int = 0
	for {
		// 返回一个时间戳 返回一个秒数： time.Now().Unix()
		// 纳秒  time.Now().UnixNano()
		rand.Seed(time.Now().UnixNano())
		n := rand.Intn(100) + 1 // [0,100)
		fmt.Println(n)
		conut++
		if n == 99 {
			break
		}
	}
	fmt.Println("生成99一共使用了", conut)
}
