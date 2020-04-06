package main

import "fmt"

func main() {
	var a []int = []int{1, 2, 3, 4, 5}
	var slice = make([]int, 10) // 10个0的一个切片
	fmt.Println(slice)

	copy(slice, a)
	fmt.Println(slice) // 将a上的元素复制到全为0的切片中
	// 切片才能进行拷贝操作
	// [0 0 0 0 0 0 0 0 0 0]
	// [1 2 3 4 5 0 0 0 0 0]

}
