package main

import "fmt"

func main() {
	var slice []int = make([]int, 4, 10)
	fmt.Println(slice)
	fmt.Println("slice len=",len(slice),"slice cap=",cap(slice))
	slice[0] = 100
	slice[2] = 200
	fmt.Println(slice)

}