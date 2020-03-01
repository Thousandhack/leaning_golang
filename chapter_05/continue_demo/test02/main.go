package main

import "fmt"

func main(){
	// continue 打印1--100内的奇数 
	for i := 1; i <= 100; i++{
		if i % 2 == 0 {
			continue
		}
		fmt.Println("i=",i)
	}
}