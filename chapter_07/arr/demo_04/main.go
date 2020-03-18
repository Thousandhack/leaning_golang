package main

import "fmt"

func test(arr [3]int) {
	arr[0] = 88
}

func demo(arr *[3]int) {
	(*arr)[1] = 99 
}

func main() {
	arr := [3]int{11,22,33}
	fmt.Println(arr)
	test(arr)
	// 数组的值不会因为调用的函数改变而改变
	// 拷贝的思想
	fmt.Println(arr)
	// demo 修改指针的方法可以实现修改数组里面的值
	// 传地址
	demo(&arr)
	fmt.Println(arr)
}

