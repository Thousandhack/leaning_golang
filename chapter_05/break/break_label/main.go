package main

import "fmt"

func main(){
	// 演示标签形式使用break
	lable1:
	for i := 0; i < 4; i++ {
		for j := 0; j < 10; j++ {
			if j == 2 {
				break lable1
			}
			fmt.Println("j=",j)  // 只打印：j = 0 j = 1 因为此时直接跳出两层循环
		}
	}
}