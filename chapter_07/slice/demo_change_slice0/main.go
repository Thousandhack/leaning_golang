package main

import "fmt"

func main() {
	var slice1 []int 
	var arr [5]int = [...]int {1,2,3,4,5} // 数组
	slice1 = arr[:]   // slice1 属于arr数组的一部分，slice2的指针和slice1存的数据一样
	var slice2 = slice1
	slice2[0] = 10
	fmt.Println(slice1)

	fmt.Println(slice2)

	fmt.Println(arr)
	// [10 2 3 4 5]
	// [10 2 3 4 5]
	// [10 2 3 4 5]

}