package main

import "fmt"

func main(){
	// continue 案例
	for i := 0; i < 4; i++ {
		if i == 2 {
			continue
		}
		fmt.Println("i=",i)  // 不打印 i = 2
	}
}