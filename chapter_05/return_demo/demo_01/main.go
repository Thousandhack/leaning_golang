package main

import "fmt"

func main() {
	// 举例return
	for i := 1; i <= 10; i++ {
		for j := 1; i <= 10; j++ {
			if j == 3 {
				return
			}
			fmt.Println(j)
		}
	}
}