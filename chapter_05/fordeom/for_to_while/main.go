package main

import "fmt"

func main() {
	// 使用for来实现do...while的方法
	var j int = 1
	for{
		fmt.Println("hello,ok",j)
		j ++ // 循环变量的迭代
		if j > 10 {
			break
		}
	}
}