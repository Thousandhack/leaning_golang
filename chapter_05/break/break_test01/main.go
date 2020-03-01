package main

import "fmt"

func main(){
	// 100以内的数求和，求和第一次超过20 时的数
	sum := 0
	for i := 1; i <= 100; i++ {
		sum += i
		if sum > 20 {
			fmt.Println(sum)
			break
		}
	}
}