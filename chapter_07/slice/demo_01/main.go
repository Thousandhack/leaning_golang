package main

import "fmt"

func main() {
	// 先定义一个数组
	var intArr [5]int = [...]int{1, 22, 33, 66, 99}

	// 定义一个切片
	slice := intArr[1:3]
	// 表示引用 intArr数组的下标为1的元素 到 下标为3 的元素，结束下标的元素不算在内
	fmt.Println(slice)       // [22 33]
	fmt.Println(len(slice))  // 2
	fmt.Println(cap(slice)) // 切片的容量是可以动态变化的，一般是slice的两倍 

	slice[1] = 44
	// 改变切片的值，会使得切片引用的数组相应的数据也改变
	fmt.Println(intArr[2]) // 44
}
