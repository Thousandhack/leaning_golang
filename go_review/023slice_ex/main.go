package main

import (
	"fmt"
	"sort"
)

func main() {
	var a = make([]int, 5, 10)
	// [0 0 0 0 0]
	for i := 0; i < 10; i++ {
		a = append(a, i)
	}
	// [0 0 0 0 0 0 1 2 3 4 5 6 7 8 9]
	fmt.Println(a)

	// 请使用内置的sort包对数组var a = [...]int{3, 7, 8, 9, 1}进行排序。
	// 对切片进行排序
	var a1 = [...]int{3, 7, 8, 9, 1}
	sort.Ints(a1[:])
	// [1 3 7 8 9]
	fmt.Println(a1)
}
