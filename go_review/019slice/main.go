package main

import "fmt"

func main() {
	// 声明一个整型的切片
	var a1 []int
	a1 = []int{1, 3, 5}

	fmt.Printf("%T\n", a1)
	fmt.Println(cap(a1))
	// 	基于数组通过切片表达式得到切片
	// 切片的容量是底层数组的容量
	a := [5]int{1, 2, 3, 4, 5}
	s := a[1:3] // s := a[low:high]
	fmt.Printf("s:%v len(s):%v cap(s):%v\n", s, len(s), cap(s))

	// 切片再切割
	s1 := s[:1]
	fmt.Println(s1)

}
