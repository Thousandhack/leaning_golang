package main

import "fmt"

func main() {
	arr := make([]int, 0)
	for i := 0; i < 10; i++ {
		fmt.Println("len 为", len(arr), "cap 为", cap(arr))
		arr = append(arr, i)
		fmt.Println(arr)
	}

}
