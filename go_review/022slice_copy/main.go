package main

import "fmt"

func main() {
	// copy 复制切片
	a1 := []int{1, 3, 7}
	a2 := a1
	var a3 = make([]int, 3, 3)
	copy(a3, a1) // 使用copy函数将切片a中的元素复制到切片c
	fmt.Println(a1)
	fmt.Println(a2)
	fmt.Println(a3)
	a1[0] = 100
	fmt.Println(a1)
	fmt.Println(a2)
	fmt.Println(a3)
	/*
		[100 3 7]
		[100 3 7]
		[1 3 7]
	*/

	// 删除 索引 为1的元素
	a3 = append(a3[:1], a3[2:]...)
	fmt.Println(a3)

	//

}
